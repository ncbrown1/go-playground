import React from 'react';

export default React.createClass({
    getInitialState() {
        return {
            showShare: false
        };
    },
    render() {
        let shareURL = '';
        if (this.state.showShare) {
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
                    <li><p className="navbar-btn"><a href="#" className="btn btn-primary" id="go-run">Run</a></p></li>
                    <li><p className="navbar-btn"><a href="#format" className="btn btn-primary" id="format">Format</a></p></li>
                    <li><label className="navbar-btn btn btn-primary">
                        <input type="checkbox" id="go-imports" />
                        {" Imports"}
                    </label></li>
                    <li><p className="navbar-btn"><a href="#share" className="btn btn-primary" id="share">Share</a></p></li>
                    <li>{ shareURL }</li>
                </ul>
                <ul className="nav navbar-nav navbar-right">
                    <li><p className="navbar-btn"><a href="#about" className="btn btn-primary" id="about">About</a></p></li>
                </ul>
            </div>
        </nav>;
    }
});