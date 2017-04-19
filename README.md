# Go Playground (Redux)

[UCSB](http://cs.ucsb.edu) - [CS 263](http://cs.ucsb.edu/~cs263) - [Chandra Krintz](http://www.cs.ucsb.edu/~ckrintz/) - Spring 2017

This is a reimplementation of the Go playground [https://play.golang.org] in Go to run via the browser (javascript + go) without the
restrictions of the original (but requiring a login to bypass those restrictions).

# Getting Started

This package uses Go and Node.JS-based packages in its implementation. Ensure you have those tools installed.

Go installation: [https://golang.org/doc/install](https://golang.org/doc/install)\
Node.JS/NPM: [https://nodejs.org/en/download/package-manager/](https://nodejs.org/en/download/package-manager/)

Note that all development for Go-related packages should take place in your `$GOPATH/src` directory. If you have not set
up your `$GOPATH`, be sure to create a new empty directory somewhere and export that directory as the value. For example:

```
$ echo $GOPATH

$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ echo $GOPATH
/Users/ncbrown/go
$ mkdir $GOPATH/src
```

Also be sure to create a directory `$GOPATH/src` to hold your go code. This is where you should be working when developing
in Go.

If you are unfamiliar with Go, you may be used to retrieving a project from github.com via `git clone`. However, the Go
ecosystem is slightly different. You will want to retrieve this project via `go get github.com/ncbrown1/go-playground`,
which will check out this repository under `$GOPATH/src/github.com/ncbrown1/go-playground`.

For package management in Go and Node.JS, we use [govendor](https://github.com/kardianos/govendor) and
[Yarn](https://yarnpkg.com), respectively. Ensure you have those tools installed.

govendor: `go get -u github.com/kardianos/govendor`\
Yarn: `npm install -g yarn`

Now that you have the proper tooling, you can update your dependency trees to gather all the required packages and dependencies:

```
$ govendor sync
$ yarn install
```

With all the required tools and packages, you can run a development version of the service via `./serve`.

To run in production, run `go install`, which will create `$GOPATH/bin/go-playground`. Run that command from this directory.

# Development

All Go source files are located in the `app/` subdirectory. All front-end resources are located in the `resources/` subdirectory. Anything in the `public/` subdirectory is directly served from the web server off of the root (`"/"`) path.
 
Style resources located in the `resources/css` directory are written in [Sass/Scss](http://sass-lang.com/). The entry point for all styles is `resources/css/app.scss`. Any styles referred to or imported in that file will be included in the final transpiled/compiled output (`public/css/main.css`).

Javascript resources located in the `resources/js` directory are written using the new syntax of javascript, i.e. ES6/ES2015. The entry point for all javascript sources is `resources/js/app.js`. Any scripts/libraries referred to or imported in that file will be included in the final transpiled/compiled output (`public/js/bundle.js`).

To add a Go dependency: `govendor fetch github.com/foo/bar`\
To add a Yarn dependency: `yarn add foo/bar`

# Project Vision

This project will consist of three parts: the web page for showing content and allowing input, the web socket for transmitting data
between the client and the server, and the job runner to compile and run the playground code on the fly.

The web page portion will be a relatively simple replica of the original at [https://play.golang.org], except the javascript functionality
will be replaced with a websocket connection to our Go server. When the user makes a request, the page's javascript will send the code
written by the user and the desired action to the server. The websocket server will respond with an output stream from the job runner.

The web socket layer is a small communication layer between web clients and the actual Go job runner, allowing requests to be made
that compile, format, and run Go code. When a request is received, this layer will send the parameters to the job runner and open an
output stream from the job runner to the web client. As this layer scans the output, if it sees a terminating character/signal, it will
close the output stream and notify the client, which will signal the end of the request. This layer is in charge of queueing requests and
load balancing to any number of Go job runners. It also does some caching to make sure that repeated requests don't take too much time.

The Go job runner will be connected to the web socket layer (not clients), processing the forwarded client requests. I believe that this
portion should just be a matter of making a call to the go binary in some sandboxed environment. As such, this could and should be simple
to implement.

I will be taking advantage of the article written about the official playground implementation here: https://blog.golang.org/playground

To take this project to the next level beyond what is written above, it may be worthwhile to investiage the use of Google's "Native
Client", which allows compiled C/C++ (and Go) code to be run in the browser efficiently and securely. In this case, the output stream
from the job runner will stream the compiled binary to the web client instead of the output from the program being run on the server.
This investigation would be very interesting, as the code would be running under different circumstances than usual. There are many
aspects of running in a browser that are quite different than running on hardware, such as the file system, the network, and concurrency.
