import React from 'react';
import { connect } from 'react-redux';
import * as actions from '../actions';

class Nav extends React.Component {
    run(e) {
        e.preventDefault();
        this.props.runCode();
    }

    fmt(e) {
        e.preventDefault();
        this.props.fmtCode();
    }

    changeLive(e) {
        this.props.switchLive(! this.props.live);
    }

    render() {
        let shareURL = '';
        if (this.props.sharing) {
            shareURL = <form className="navbar-form">
                <input type="text" id="share-url" className="form-control" value="https://bit.ly/asdsfafsd"/>
            </form>;
        }
        return <nav className="navbar navbar-default">
            <div style={{paddingRight: "2em"} /* fixing overflow to the right */ }>
                <div className="navbar-header">
                    <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false"      aria-controls="navbar">
                        <span className="sr-only">Toggle navigation</span>
                        <span className="icon-bar"/>
                        <span className="icon-bar"/>
                        <span className="icon-bar"/>
                    </button>
                    <a className="navbar-brand logo" href="#">The Go Playground (Redux)</a>
                </div>
                <ul className="nav navbar-nav">
                    <li><p className="navbar-btn">
                        <a href="#" className="btn btn-primary" id="go-run" onClick={this.run.bind(this)}>Run</a>
                    </p></li>
                    <li><p className="navbar-btn">
                        <a href="#format" className="btn btn-primary" id="format" onClick={this.fmt.bind(this)}>Format</a>
                    </p></li>
                    <li><label className="navbar-btn btn btn-primary">
                        <input type="checkbox" id="go-imports" />
                        {" Imports"}
                    </label></li>
                    <li><p className="navbar-btn"><a href="#share" className="btn btn-primary" id="share">Share</a></p></li>
                    <li>{ shareURL }</li>
                </ul>
                <ul className="nav navbar-nav navbar-right">
                    <li><label className="navbar-btn btn btn-primary">
                        <input type="checkbox" id="go-live" checked={this.props.live} onChange={this.changeLive.bind(this)} />
                        {" WebSockets"}
                    </label></li>
                    <li><p className="navbar-btn"><a href="https://github.com/ncbrown1/go-playground" className="btn btn-primary" id="about">About</a></p></li>
                </ul>
            </div>
        </nav>;
    }
}

// Maps state from store to props
const mapStateToProps = (state, ownProps) => {
    return {
        sharing: state.sharing, // this.props.sharing
        live: state.live
    };
};

// Maps actions to props
const mapDispatchToProps = (dispatch) => {
    return {
        runCode: () => dispatch(actions.runCode()),
        fmtCode: () => dispatch(actions.fmtCode()),
        switchLive: (live) => dispatch(actions.switchLive(live))
    };
};

// Use connect to put them together
export default connect(mapStateToProps, mapDispatchToProps)(Nav);