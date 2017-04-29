import React from 'react';
import CodeMirror from 'react-codemirror';

const initial_program = `
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, playground")
}
`.trim();

export default React.createClass({
    getInitialState() {
        return {
            code: initial_program
        };
    },
    updateCode(newCode) {
        this.setState({
            code: newCode
        });
    },
    render() {
        let options = {
            autofocus: true,
            indentUnit: 4,
            keymap: 'sublime',
            lineNumbers: true,
            mode: 'go',
            theme: 'monokai',
            value: initial_program
        };
        return <div id="go-editor">
            <CodeMirror value={this.state.code} onChange={this.updateCode} options={options} />
        </div>;
    }
})