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
)

func RunCode(code string) (*Result) {
    h := sha1.New()
    io.WriteString(h, code)
    sha := fmt.Sprintf("%x", h.Sum(nil))

    filename := fmt.Sprintf("/tmp/%s.go", sha)
    err := ioutil.WriteFile(filename, []byte(code), 0644)
    if err != nil {
        panic(err)
    }

    result := runGoProgram(filename)
    return &result
}

func runGoProgram(file string) Result {
    pid := os.Getpid()
    exec_file := fmt.Sprintf("/tmp/prog-%d", pid)
    compile_file := fmt.Sprintf("%s-compiled.o", exec_file)
    compiled := run("go", "tool", "compile", "-o", compile_file, file)
    run("rm", compile_file) // cleanup unneeded file

    if compiled.success() {
        run("go", "build", "-o", exec_file, file)
        result := run(exec_file)

        run("rm", exec_file) // cleanup unneeded file
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