package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola desde el backend en Go!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/api/hello", helloHandler)

    fmt.Println("Servidor iniciado en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}