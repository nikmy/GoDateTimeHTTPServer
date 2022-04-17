package main

import (
    "log"
    "net/http"
    "time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    _, err := w.Write([]byte("<h1>" + time.Now().Format(time.UnixDate) + "</h1>"))
    if err != nil {
        log.Println(err)
    }
}

func main() {
    http.HandleFunc("/", handleRequest)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal(err)
    }
}
