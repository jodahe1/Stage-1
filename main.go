package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/name", nameHandler)
	http.HandleFunc("/hobby", hobbyHandler)
	http.HandleFunc("/dream", dreamHandler)

	fmt.Println("Server is starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
