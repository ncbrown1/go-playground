import React from 'react';

export default React.createClass({
    getInitialState() {
        return {
            stdout: 'Hello, playground.',
            system: 'Program exited.'
        }
    },
    render() {
        return <div id="go-output">
            <pre>
                <span className="stdout">{this.state.stdout}</span>
                <br/>
                <span className="system">{this.state.system}</span>
            </pre>
        </div>;
    }
});