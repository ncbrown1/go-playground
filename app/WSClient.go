package app

import (
    "net/http"
    "fmt"

    "github.com/gorilla/websocket"
    "github.com/ncbrown1/go-playground/app/runtime"
)

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    runtime.RunSocket(conn)
}