package main

import (
    "os/exec"
    "fmt"
    "flag"
)

func main() {
    filePtr := flag.String("file", "", "file path")
    flag.Parse()

    if len(*filePtr) > 0 {
        fmt.Print("File chosen: ")
        fmt.Println(*filePtr)

        cmd := exec.Command("go", "run", *filePtr)
        output, err := cmd.Output()

        fmt.Println("stdout:")
        fmt.Print(string(output[:]))

        fmt.Println("\nstderr:")
        if err == nil {
            fmt.Println("NULL")
        } else {
            fmt.Println(err)
        }
    } else {
        fmt.Println("YOU MUST CHOOSE A PROPER FILE TO RUN")
    }
}