package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func nameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Yodahe Teshome")) 
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}


type Hobby struct {
	FavoriteHobby string `json:"favorite_hobby"`
}


func hobbyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		hobby := Hobby{
			FavoriteHobby: "Playing and Watching Football , Reading Books", 
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hobby)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}



type Dream struct {
	Message string `json:"message"`
}


func dreamHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		dream := Dream{
			Message: "I aim to solve real-world problems, and create innovative solutions that positively impact communities and businesses",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dream)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/name", nameHandler)
	http.HandleFunc("/hobby", hobbyHandler)
	http.HandleFunc("/dream", dreamHandler)

	fmt.Println("Server is starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}