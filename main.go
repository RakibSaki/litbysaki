package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    port := os.Getenv("PORT")
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
