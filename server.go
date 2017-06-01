package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/ncbrown1/go-playground/app"
)

func main() {
    var port, address string

    flag.StringVar(&port, "p", "8000", "the port to bind to for serving")
    flag.StringVar(&port, "port", "8000", "the port to bind to for serving")
    flag.Parse()

    address = fmt.Sprintf(":%s", port)

    router := gin.New()
    router.Use(gin.Logger())

    router = app.SetupRoutes(router)


    // This will serve all public files under /<filename>
    router.StaticFS("/", http.Dir("./public/"))

    log.Printf("Serving site in port %s\n", address)

    srv := &http.Server{
        Addr:           address,
        Handler:        router,
        ReadTimeout:    15 * time.Second,
        WriteTimeout:   15 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    log.Fatal(srv.ListenAndServe())
}
