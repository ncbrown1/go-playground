
require('./setup');

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
