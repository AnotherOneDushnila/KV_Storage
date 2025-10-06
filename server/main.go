package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/AnotherOneDushnila/KV_Storage/internal/store"
	httpapi "github.com/AnotherOneDushnila/KV_Storage/internal/api/http"
	// badger "github.com/AnotherOneDushnila/KV_Storage/internal/store/badger"
	"github.com/AnotherOneDushnila/KV_Storage/internal/store/inmemory"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	backend := os.Getenv("STORAGE_BACKEND")
	fmt.Println(backend)

	
	s := memory.New()
	handler := httpapi.NewHandler(s)
	http.HandleFunc("/ping", handler.Ping)
	http.HandleFunc("/put", handler.PutHandler)
	http.HandleFunc("/get", handler.GetHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)


	log.Println("Starting storage on :8080")
	http.ListenAndServe(":8080", nil)
}