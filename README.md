# Go Playground (Redux)

[UCSB](http://cs.ucsb.edu) - [CS 263](http://cs.ucsb.edu/~cs263) - [Chandra Krintz](http://www.cs.ucsb.edu/~ckrintz/) - Spring 2017

This is a reimplementation of the Go playground [https://play.golang.org] in Go to run via the browser (javascript + go) without the
restrictions of the original (but requiring a login to bypass those restrictions).

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
