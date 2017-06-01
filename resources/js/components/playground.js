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
            stdout: '',
            system: ''
        };
    },
    updateCode(newCode) {
        this.setState({
            code: newCode
        });
    },
    runCode(e) {
        e.preventDefault();
        this.setState({
            stdout: 'Waiting for remote server...',
            system: ''
        });
        axios
            .post('/run-code', {code: this.state.code})
            .then((resp) => {
                console.log(resp.data);
                if (resp.data.hasOwnProperty('output')) {
                    this.setState({
                        stdout: resp.data.output,
                        system: 'Program exited.'
                    });
                }
            })
            .catch((err) => {
                console.log(err);
            });
    },
    fmtCode(e) {
        e.preventDefault();
        console.log('Formatting Code');
    },
    shouldComponentUpdate(nextProps, nextState) {
        let old_out = this.state.stdout, old_sys = this.state.system,
            new_out = nextState.stdout, new_sys = nextState.system;
        return !(old_out === new_out || old_sys === new_sys);
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