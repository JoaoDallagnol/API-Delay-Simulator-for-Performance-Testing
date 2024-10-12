package main

import (
	"log"
	"net/http"

	"github.com/JoaoDallagnol/API-Delay-Simulator-for-Performance-Testing.git/internal/handler"
)

func main() {
	http.HandleFunc("/delaye", handler.DelayHandler)
	http.HandleFunc("/unstable", handler.UnstableHandler)
	http.HandleFunc("/custom-delay", handler.CustomDelayHandler)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Falied to start server: %v", err)
	}
}
