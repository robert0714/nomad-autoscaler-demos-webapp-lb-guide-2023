package main

import (
        "fmt"
        "net/http"
        "os"
)

// Environment variables NODE_IP and PORT must be set or passed in

var host string = os.Getenv("NODE_IP")
var port string = os.Getenv("PORT")

func giveIntro(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome! You are on node %v:%v\n", host, port)
}

func main() {
        mux := http.NewServeMux()
        server := &http.Server{
                Addr:    fmt.Sprintf("0.0.0.0:%v", port),
                Handler: mux,
        }
        mux.HandleFunc("/", giveIntro)
        server.ListenAndServe()
}
