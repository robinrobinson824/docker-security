package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `<html>
        <head><title>Hello</title></head>
        <body style="background-color: purple; display: flex; align-items: center; justify-content: center; height: 100vh;">
            <h1 style="color: white;">Hello, World!</h1>
        </body>
    </html>`)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

