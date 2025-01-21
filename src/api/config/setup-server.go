package config

import (
	"log"
	"net/http"
)

func InitServer() {
	log.Println("Starting our simple http server.")
	port := ":8080"
	log.Printf("Started on http://localhost%s", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
