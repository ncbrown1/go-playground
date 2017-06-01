package app

import (
    "net/http"
    "encoding/json"
    "log"
)

type RunCodeJSON struct {
    Code string
}

func Index(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w,"index.html", nil)
}

func RunCode(w http.ResponseWriter, r *http.Request) {
    var run_code RunCodeJSON
    if r.Body == nil {
        http.Error(w, "Please send a request body", 400)
        return
    }
    err := json.NewDecoder(r.Body).Decode(&run_code)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }
    log.Printf("Program:\n%s", run_code.Code)
    w.WriteHeader(204)
}