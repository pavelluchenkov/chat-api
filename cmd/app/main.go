package main

import (
	"chat-api/internal/db"
	"chat-api/internal/models"
	"chat-api/internal/router"
	"log"
	"net/http"
)

func main() {

	database := db.Connect()

	
	if err := database.AutoMigrate(&models.Chat{}, &models.Message{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	r := router.NewRouter(database)
	
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
