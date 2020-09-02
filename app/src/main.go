package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, greeting("CodeEducation Rocks!"))
}

func greeting(message string) string {
    return "<b>" + message + "</b>"
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}