import React from 'react';

export default React.createClass({
    render() {
        return <div id="go-output">
            <pre>
                <span className="stdout">{this.props.stdout}</span>
                <br/>
                <span className="system">{this.props.system}</span>
            </pre>
        </div>;
    }
});