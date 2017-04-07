package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var sdir, host, port, address string

	flag.StringVar(&sdir, "dir", "./static/", "the directory to serve static files from")
	flag.StringVar(&host, "host", "127.0.0.1", "the host to bind to for serving")
	flag.StringVar(&port, "port", "8000", "the port to bind to for servinbg")
	flag.Parse()
	router := mux.NewRouter()

	address = fmt.Sprintf("%s:%s", host, port)

	// This will serve files under http://localhost:8000/static/<filename>
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(sdir))))
	router.HandleFunc("/", RootHandler)

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

func RootHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}
