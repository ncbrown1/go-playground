import React from 'react';
import Nav from './nav';
import Editor from './editor';
import Output from './output';

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
            code: initial_program,
            stdout: 'Hello, playground.',
            system: 'Program exited.'
        };
    },
    updateCode(newCode) {
        this.setState({
            code: newCode
        });
    },
    runCode(e) {
        e.preventDefault();
        console.log("Running Code");
        console.log(this.state.code);
    },
    fmtCode(e) {
        e.preventDefault();
        console.log("Formatting Code");
    },
    shouldComponentUpdate(nextProps, nextState) {
        return !nextState.hasOwnProperty("code");

    },
    render() {
        return <div>
            <Nav runCode={this.runCode}
                 fmtCode={this.fmtCode} />
            <Editor code={this.state.code}
                    updateCode={this.updateCode}
                    key="go-editor" />
            <Output stdout={this.state.stdout}
                    system={this.state.system} />
        </div>;
    }
})