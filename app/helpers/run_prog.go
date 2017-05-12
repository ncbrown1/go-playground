package main

import (
    "os"
    "os/exec"
    "fmt"
    "flag"
    "syscall"
    "bytes"
)

type Result struct {
    code   int
    output string
    error  string
}

func main() {
    filePtr := flag.String("file", "", "file path")
    flag.Parse()

    if len(*filePtr) > 0 {
        fmt.Print("File chosen: ")
        fmt.Println(*filePtr)

        pid := os.Getpid()
        exec_file := fmt.Sprintf("/tmp/prog-%d", pid)
        compile_file := fmt.Sprintf("%s-compiled.o", exec_file)
        compiled := run("go", "tool", "compile", "-o", compile_file, *filePtr)

        if compiled.code == 0 {
            fmt.Printf("Successfully compiled to: %s\n", exec_file)

            run("go", "build", "-o", exec_file, *filePtr)
            execd := run(exec_file)

            fmt.Println("stdout:")
            fmt.Println(execd.output)

            fmt.Println("stderr:")
            fmt.Println(execd.error)

            fmt.Printf("Exit status: %d\n", execd.code)
        } else {
            fmt.Print(compiled.output)
        }
    } else {
        fmt.Println("YOU MUST CHOOSE A PROPER FILE TO RUN")
    }
}

func run(exe string, args ...string) Result {
    var exit_code int
    cmd := exec.Command(exe, args...)
    cmdOutput := &bytes.Buffer{}
    cmdError := &bytes.Buffer{}
    cmd.Stdout = cmdOutput
    cmd.Stderr = cmdError

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
        string(cmdError.Bytes()),
    }
}