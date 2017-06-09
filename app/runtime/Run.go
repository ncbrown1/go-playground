package runtime

import (
    "crypto/sha1"
    "io"
    "fmt"
    "os/exec"
    "bytes"
    "syscall"
    "os"
    "io/ioutil"
    "log"
    "encoding/json"
    "time"
    //"path/filepath"
    "github.com/gorilla/websocket"
)

//// RunScripts specifies whether the socket handler should execute shell scripts
//// (snippets that start with a shebang).
//var RunScripts = true

// Environ provides an environment when a binary, such as the go tool, is
// invoked.
var Environ func() []string = os.Environ

const (
    // The maximum number of messages to send per session (avoid flooding).
    msgLimit = 1000

    // Batch messages sent in this interval and send as a single message.
    msgDelay = 10 * time.Millisecond
)

func RunCode(code string) (*Result) {
    h := sha1.New()
    io.WriteString(h, code)
    sha := fmt.Sprintf("%x", h.Sum(nil))

    out_filename := fmt.Sprintf("/tmp/%s.output", sha)
    src_filename := fmt.Sprintf("/tmp/%s.go", sha)
    var b []byte
    var result Result
    var err error

    // read a Result from the file and if no error, try to decode
    if b, err = ioutil.ReadFile(out_filename); err == nil {
        result, err = decodeResults(b)
    }

    // if there was an error, compile and run the code again
    if err != nil {
        log.Println("Compiling and running code")
        err := ioutil.WriteFile(src_filename, []byte(code), 0644)
        defer os.Remove(src_filename)
        if err != nil {
            panic(err)
        }
        result = runGoProgram(src_filename)
        ioutil.WriteFile(out_filename, result.encode(), 0644)
    }

    return &result
}

func runGoProgram(file string) Result {
    pid := os.Getpid()
    exec_file := fmt.Sprintf("/tmp/prog-%d", pid)
    compile_file := fmt.Sprintf("%s-compiled.o", exec_file)
    compiled := run("go", "tool", "compile", "-o", compile_file, file)
    defer os.Remove(compile_file) // cleanup unneeded file

    if compiled.success() {
        run("go", "build", "-o", exec_file, file)
        result := run(exec_file)

        defer os.Remove(exec_file) // cleanup unneeded file
        return result
    } else {
        return compiled
    }
}

func run(exe string, args ...string) Result {
    var exit_code int
    cmd := exec.Command(exe, args...)
    cmdOutput := &bytes.Buffer{}
    cmd.Stdout = cmdOutput
    cmd.Stderr = cmdOutput

    if err := cmd.Run(); err != nil {
        //if err != nil {
        //    os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
        //}
        if exitError, ok := err.(*exec.ExitError); ok {
            exit_code = exitError.Sys().(syscall.WaitStatus).ExitStatus()
        }
    } else {
        // Success
        exit_code = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
    }

    return Result{
        exit_code,
        string(cmdOutput.Bytes()),
    }
}

//----------------------------------------------------------------------------------------------------------------------

func RunSocket(c *websocket.Conn) {
    in, out := make(chan *Message), make(chan *Message)
    errc := make(chan error, 1)

    // Decode messages from client and send to the in channel.
    go func() {
        for {
            _, r, err := c.NextReader()
            if err != nil {
                log.Println(err)
            }
            dec := json.NewDecoder(r)
            var m Message
            if err := dec.Decode(&m); err != nil {
                errc <- err
                return
            }
            in <- &m
        }
    }()

    // Receive messages from the out channel and encode to the client.
    go func() {
        for m := range out {
            w, err := c.NextWriter(websocket.TextMessage)
            if err != nil {
                log.Println(err)
            }
            enc := json.NewEncoder(w)
            if err := enc.Encode(m); err != nil {
                errc <- err
                return
            }
        }
    }()
    defer close(out)

    // Start and kill Processes and handle errors.
    proc := make(map[string]*Process)
    for {
        select {
        case m := <-in:
            switch m.Kind {
            case "run":
                log.Println("running snippet from:", c.RemoteAddr().String())
                proc[m.Id].Kill()
                proc[m.Id] = StartProcess(m.Id, m.Body, out)
            case "kill":
                proc[m.Id].Kill()
            }
        case err := <-errc:
            // A encode or decode has failed; bail.
            log.Println(err)
            // Shut down any running processes.
            for _, p := range proc {
                p.Kill()
            }
            return
        }
    }
}