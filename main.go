package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Golang! - MD WASIM ANSARI")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 8081...")
    http.ListenAndServe(":8081", nil)
}

