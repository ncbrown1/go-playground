package main

import (
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w,"index.html", nil)
}
