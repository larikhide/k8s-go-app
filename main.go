package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler (w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func main() {
	http.HandleFunc("/", handler)

	port := "8080"
	log.Printf("start server on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}