package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    instanceID := os.Getenv("INSTANCE_ID")
    port := os.Getenv("PORT")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "HELLO FROM INSTANCE: %s ON PORT %s\n", instanceID, port)
    })

    fmt.Println("INSTANCE ",instanceID, "running on port", port)
    http.ListenAndServe(":"+port, nil)
}
