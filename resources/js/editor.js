const initial_program = `
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, playground")
}
`.trim();

CodeMirror(document.getElementById("go-editor"), {
    autofocus: true,
    indentUnit: 4,
    keymap: 'sublime',
    lineNumbers: true,
    mode: 'go',
    theme: 'monokai',
    value: initial_program
});

const sample_output = 'Hello, playground';
const sample_system = 'Program exited.';

let stdout = $("#go-output .stdout");
let sysout = $("#go-output .system");

if (stdout.length) stdout.innerText = sample_output;
if (sysout.length) sysout.innerText = sample_system;