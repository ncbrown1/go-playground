package app

import (
    "net/http"
    "fmt"

    "github.com/gorilla/websocket"
    "github.com/ncbrown1/go-playground/app/runtime"
)

// Message is the wire format for the websocket connection to the browser.
// It is used for both sending output messages and receiving commands, as
// distinguished by the Kind field.
type Message struct {
    Id   string `json:"id"`   // client-provided unique id for the process
    Kind string `json:"kind"` // in: "run", "kill" out: "stdout", "stderr", "end"
    Body string `json:"body"`
}

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