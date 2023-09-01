package main

import (
    "html/template"
    "net/http"
)

func main() {

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Create a template for the HTML page.
    tmpl := template.Must(template.ParseFiles("index.html"))

    // Handle requests for the index.html page.
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    http.ListenAndServe(":8080", nil)
}

