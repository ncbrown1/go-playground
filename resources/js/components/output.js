import React from 'react';
import { connect } from 'react-redux';

class Output extends React.Component {
    render() {
        return <div id="go-output">
            <pre>
                <span className="stdout">{this.props.output}</span>
                {/*<br/>*/}
                <span className="system">{this.props.system}</span>
            </pre>
        </div>;
    }
};

// Maps state from store to props
const mapStateToProps = (state, ownProps) => {
    return {
        output: state.output, // this.props.output
        system: state.system, // this.props.output
    };
};

// Maps actions to props
const mapDispatchToProps = (dispatch) => {
    return {};
};

// Use connect to put them together
export default connect(mapStateToProps, mapDispatchToProps)(Output);