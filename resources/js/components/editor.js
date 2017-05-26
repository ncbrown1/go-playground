import React from 'react';
import CodeMirror from 'react-codemirror';

export default React.createClass({
    render() {
        let options = {
            autofocus: true,
            indentUnit: 4,
            keymap: 'sublime',
            lineNumbers: true,
            mode: 'go',
            theme: 'monokai',
            value: this.props.code
        };
        return <div id="go-editor">
            <CodeMirror value={this.props.code} onChange={this.props.updateCode} options={options} />
        </div>;
    }
})