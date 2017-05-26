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
    render() {
        return <div>
            <Nav/>
            <Editor code={this.state.code} updateCode={this.updateCode} />
            <Output stdout={this.state.stdout} system={this.state.system} />
        </div>;
    }
})