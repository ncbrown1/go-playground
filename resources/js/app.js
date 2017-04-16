
require('./setup');

function resizeInput() {
    $(this).attr('size', $(this).val().length);
}

let share_url = $('#share-url');
share_url.keyup(resizeInput)
    // resize on page load
    .each(resizeInput);
$('#share').click(() => {
    share_url.addClass('active');
    share_url.focus();
});

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

document.querySelector("#go-output .stdout").innerText = sample_output;
document.querySelector("#go-output .system").innerText = sample_system;