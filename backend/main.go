package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from Go!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", helloHandler)

	c := cors.Default()
	handler := c.Handler(mux)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
