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

        result := runGoProgram(*filePtr)
        fmt.Println("stdout:")
        fmt.Println(result.output)

        fmt.Println("stderr:")
        fmt.Println(result.error)

        fmt.Printf("Exit status: %d\n", result.code)
    } else {
        fmt.Println("YOU MUST CHOOSE A PROPER FILE TO RUN")
    }
}

func runGoProgram(file string) Result {
    pid := os.Getpid()
    exec_file := fmt.Sprintf("/tmp/prog-%d", pid)
    compile_file := fmt.Sprintf("%s-compiled.o", exec_file)
    compiled := run("go", "tool", "compile", "-o", compile_file, file)
    run("rm", compile_file) // cleanup unneeded file

    if compiled.code == 0 {
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