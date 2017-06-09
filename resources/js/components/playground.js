import React from 'react';
import Nav from './nav';
import Editor from './editor';
import Output from './output';

class Playground extends React.Component {
    render() {
        return <div>
            <Nav />
            <Editor />
            <Output />
        </div>;
    }
}

export default Playground;