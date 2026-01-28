package router

import (
	"chat-api/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/chats", handlers.CreateChat(db))
	r.Post("/chats/{id}/messages", handlers.CreateMessage(db))
	r.Get("/chats/{id}", handlers.GetChat(db))
	r.Delete("/chats/{id}", handlers.DeleteChat(db))

	return r
}
