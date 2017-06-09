import React from 'react';
import { connect } from 'react-redux';
import CodeMirror from 'react-codemirror';
import * as actions from '../actions';

class Editor extends React.Component {
    constructor(props){
        super(props);
    }

    shouldComponentUpdate(nextProps, newState) {
        return nextProps.forceEditorRender;
    }

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
}

// Maps state from store to props
const mapStateToProps = (state, ownProps) => {
    return {
        code: state.code, // this.props.code
        forceEditorRender: state.forceEditorRender
    };
};

// Maps actions to props
const mapDispatchToProps = (dispatch) => {
    return {
        // this.props.updateCode
        updateCode: code => dispatch(actions.updateCode(code))
    };
};

// Use connect to put them together
export default connect(mapStateToProps, mapDispatchToProps)(Editor);