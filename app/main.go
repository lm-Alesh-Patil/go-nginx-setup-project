package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "Go App"
	}
	fmt.Fprintf(w, "Hello from %s!", appName)
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Println("Go app running on port", port)
	http.ListenAndServe(":"+port, nil)
}
