package main

import (
    "net/http"
    "path/filepath"
    "fmt"
    "html/template"
)

func renderTemplate(w http.ResponseWriter, tname string, data interface{}) {
    layout_file := filepath.Join("resources/views", "layout.html")
    filename := filepath.Join("resources/views", filepath.Clean(tname))

    t, err := template.ParseFiles(layout_file, filename)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "404 Not Found")
    } else {
        t.ExecuteTemplate(w, "layout", data)
    }
}