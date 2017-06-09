package runtime

// note: copied from https://github.com/golang/tools/blob/master/playground/socket/socket.go and annotated

import (
    "os/exec"
    "path/filepath"
    "strconv"
    "os"
    "io/ioutil"
    goruntime "runtime"
    "time"
    "errors"
    "strings"
    "go/parser"
    "go/token"
    "unicode/utf8"
    "bytes"
    "log"
)

// Message is the wire format for the websocket connection to the browser.
// It is used for both sending output messages and receiving commands, as
// distinguished by the Kind field.
type Message struct {
    Id   string `json:"id"`   // client-provided unique id for the process
    Kind string `json:"kind"` // in: "run", "kill" out: "stdout", "stderr", "end"
    Body string `json:"body"`
}


// Process represents a running process.
type Process struct {
    id   string
    out  chan<- *Message
    done chan struct{} // closed when wait completes
    run  *exec.Cmd
    bin  string  // compiled binary location
}

// StartProcess builds and runs the given program, sending its output
// and end event as Messages on the provided channel.
func StartProcess(id, body string, dest chan<- *Message) *Process {
    var (
        done = make(chan struct{})
        out  = make(chan *Message)
        p    = &Process{out: out, done: done}
    )
    go func() {
        defer close(done)
        // for each message in the buffer of messages
        for m := range buffer(limiter(out, p), time.After) {
            m.Id = id
            dest <- m
        }
    }()
    // start the process
    if err := p.start(body); err != nil {
        p.end(err) // if starting failed, then stop
        return nil
    }
    go func() {
        // wait until the process exits
        p.end(p.run.Wait())
    }()
    return p
}

// A killer provides a mechanism to terminate a process.
// The Kill method returns only once the process has exited.
type killer interface {
    Kill()
}

// limiter returns a channel that wraps the given channel.
// It receives Messages from the given channel and sends them to the returned
// channel until it passes msgLimit messages, at which point it will kill the
// process and pass only the "end" message.
// When the given channel is closed, or when the "end" message is received,
// it closes the returned channel.
func limiter(in <-chan *Message, p killer) <-chan *Message {
    out := make(chan *Message)
    go func() {
        // close the out channel once goroutine exits
        defer close(out)
        n := 0
        // for each message received until no more
        for m := range in {
            switch {
            case n < msgLimit || m.Kind == "end":
                out <- m
                if m.Kind == "end" {
                    return
                }
            case n == msgLimit:
                // Kill in a goroutine as Kill will not return
                // until the process' output has been
                // processed, and we're doing that in this loop.
                go p.Kill()
            default:
                continue // don't increment
            }
            n++
        }
    }()
    return out
}

// buffer returns a channel that wraps the given channel. It receives messages
// from the given channel and sends them to the returned channel.
// Message bodies are gathered over the period msgDelay and coalesced into a
// single Message before they are passed on. Messages of the same kind are
// coalesced; when a message of a different kind is received, any buffered
// messages are flushed. When the given channel is closed, buffer flushes the
// remaining buffered messages and closes the returned channel.
// The timeAfter func should be time.After. It exists for testing.
func buffer(in <-chan *Message, timeAfter func(time.Duration) <-chan time.Time) <-chan *Message {
    out := make(chan *Message)
    go func() {
        defer close(out) // close out channel once goroutine exits
        var (
            tc    <-chan time.Time
            buf   []byte
            kind  string
            flush = func() { // send messages in buffer and refresh
                if len(buf) == 0 {
                    return
                }
                out <- &Message{Kind: kind, Body: safeString(buf)}
                buf = buf[:0] // recycle buffer
                kind = ""
            }
        )
        for {
            select {
            // get a queued outgoing message
            case m, ok := <-in:
                if !ok { // flush and exit on errors
                    flush()
                    return
                }
                if m.Kind == "end" { // flush and exit on end
                    flush()
                    out <- m
                    return
                }
                if kind != m.Kind { // flush mistyped messages
                    flush()
                    kind = m.Kind
                    if tc == nil {
                        tc = timeAfter(msgDelay)
                    }
                }
                buf = append(buf, m.Body...)
            case <-tc: // flush on timeout
                flush()
                tc = nil
            }
        }
    }()
    return out
}

// startProcess starts a given program given its path and passing the given body
// to the command standard input.
func (p *Process) startProcess(path string, args []string, body string) error {
    cmd := &exec.Cmd{
        Path:   path,
        Args:   args,
        Stdin:  strings.NewReader(body),
        Stdout: &messageWriter{kind: "stdout", out: p.out},
        Stderr: &messageWriter{kind: "stderr", out: p.out},
    }
    if err := cmd.Start(); err != nil {
        return err
    }
    p.run = cmd
    return nil
}

// start builds and starts the given program, sending its output to p.out,
// and stores the running *exec.Cmd in the run field.
func (p *Process) start(body string) error {
    // We "go build" and then exec the binary so that the
    // resultant *exec.Cmd is a handle to the user's program
    // (rather than the go tool process).
    // This makes Kill work.

    bin := filepath.Join(tmpdir, "compile"+strconv.Itoa(<-uniq))
    src := bin + ".go"
    //if goruntime.GOOS == "windows" {
    //    bin += ".exe"
    //}

    // write body to x.go
    defer os.Remove(src) // remove source file on method exit
    err := ioutil.WriteFile(src, []byte(body), 0666)
    if err != nil {
        return err
    }

    // build x.go, creating x
    p.bin = bin // to be removed by p.end
    dir, file := filepath.Split(src)
    args := []string{"go", "build", "-tags", "OMIT"}
    args = append(args, "-o", bin, file)
    cmd := p.cmd(dir, args...)
    cmd.Stdout = cmd.Stderr // send compiler output to stderr
    // compile the user program
    if err := cmd.Run(); err != nil {
        return err
    }

    // run x
    if isNacl() {
        cmd, err = p.naclCmd(bin)
        if err != nil {
            return err
        }
    } else {
        cmd = p.cmd("", bin)
    }
    if err := cmd.Start(); err != nil {
        // If we failed to exec, that might be because they built
        // a non-main package instead of an executable.
        // Check and report that.
        if name, err := packageName(body); err == nil && name != "main" {
            return errors.New(`executable programs must use "package main"`)
        }
        return err
    }
    p.run = cmd
    return nil
}

// cmd builds an *exec.Cmd that writes its standard output and error to the
// Process' output channel.
func (p *Process) cmd(dir string, args ...string) *exec.Cmd {
    cmd := exec.Command(args[0], args[1:]...)
    cmd.Dir = dir
    cmd.Env = Environ()
    cmd.Stdout = &messageWriter{p.id, "stdout", p.out}
    cmd.Stderr = &messageWriter{p.id, "stderr", p.out}
    return cmd
}

// wait waits for the running process to complete
// and sends its error state to the client.
func (p *Process) wait() {
    defer close(p.done)
    p.end(p.run.Wait())
}

// end sends an "end" message to the client, containing the process id and the
// given error value.
func (p *Process) end(err error) {
    if p.bin != "" {
        defer os.Remove(p.bin)
    }
    m := &Message{Kind: "end"}
    if err != nil {
        m.Body = err.Error()
    }
    p.out <- m
    close(p.out)
}

// Kill the process if it is running and waits for it to exit.
func (p *Process) Kill() {
    if p == nil || p.run == nil {
        return
    }
    p.run.Process.Kill()
    <-p.done
}

func isNacl() bool {
    for _, v := range append(Environ(), os.Environ()...) {
        if v == "GOOS=nacl" {
            return true
        }
    }
    return false
}

// naclCmd returns an *exec.Cmd that executes bin under native client.
func (p *Process) naclCmd(bin string) (*exec.Cmd, error) {
    pwd, err := os.Getwd()
    if err != nil {
        return nil, err
    }
    var args []string
    env := []string{
        "NACLENV_GOOS=" + goruntime.GOOS,
        "NACLENV_GOROOT=/go",
        "NACLENV_NACLPWD=" + strings.Replace(pwd, goruntime.GOROOT(), "/go", 1),
    }
    switch goruntime.GOARCH {
    case "amd64":
        env = append(env, "NACLENV_GOARCH=amd64p32")
        args = []string{"sel_ldr_x86_64"}
    case "386":
        env = append(env, "NACLENV_GOARCH=386")
        args = []string{"sel_ldr_x86_32"}
    case "arm":
        env = append(env, "NACLENV_GOARCH=arm")
        selLdr, err := exec.LookPath("sel_ldr_arm")
        if err != nil {
            return nil, err
        }
        args = []string{"nacl_helper_bootstrap_arm", selLdr, "--reserved_at_zero=0xXXXXXXXXXXXXXXXX"}
    default:
        return nil, errors.New("native client does not support GOARCH=" + goruntime.GOARCH)
    }

    cmd := p.cmd("", append(args, "-l", "/dev/null", "-S", "-e", bin)...)
    cmd.Env = append(cmd.Env, env...)

    return cmd, nil
}

func packageName(body string) (string, error) {
    f, err := parser.ParseFile(token.NewFileSet(), "prog.go",
        strings.NewReader(body), parser.PackageClauseOnly)
    if err != nil {
        return "", err
    }
    return f.Name.String(), nil
}

// messageWriter is an io.Writer that converts all writes to Message sends on
// the out channel with the specified id and kind.
type messageWriter struct {
    id, kind string
    out      chan<- *Message
}

func (w *messageWriter) Write(b []byte) (n int, err error) {
    w.out <- &Message{Kind: w.kind, Body: safeString(b)}
    return len(b), nil
}

// safeString returns b as a valid UTF-8 string.
func safeString(b []byte) string {
    if utf8.Valid(b) {
        return string(b)
    }
    var buf bytes.Buffer
    for len(b) > 0 {
        r, size := utf8.DecodeRune(b)
        b = b[size:]
        buf.WriteRune(r)
    }
    return buf.String()
}

var tmpdir string

func init() {
    // find real path to temporary directory
    var err error
    tmpdir, err = filepath.EvalSymlinks(os.TempDir())
    if err != nil {
        log.Fatal(err)
    }
}

var uniq = make(chan int) // a source of numbers for naming temporary files

func init() {
    go func() {
        for i := 0; ; i++ {
            uniq <- i
        }
    }()
}