package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ncbrown1/go-playground/app"
)

func main() {
	var host, port, address string

	flag.StringVar(&host, "host", "127.0.0.1", "the host to bind to for serving")
	flag.StringVar(&port, "port", "8000", "the port to bind to for servinbg")
	flag.Parse()

	address = fmt.Sprintf("%s:%s", host, port)

	router := app.NewRouter()
	// This will serve all public files under /<filename>
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Printf("Serving site at http://%s\n", address)

	srv := &http.Server{
		Handler: router,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
