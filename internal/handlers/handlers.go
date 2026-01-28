package handlers

import (
	"chat-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateChat(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err !=nil{
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		if len(req.Title) == 0 || len(req.Title) > 200{
			http.Error(w, "title must be 1-200 chars", http.StatusBadRequest)
			return
		}
		chat := models.Chat{
			Title: req.Title,
			CreatedAt: time.Now(),
		}
		if err := db.Create(&chat).Error; err != nil {
			http.Error(w, "failed to create chat", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chat)
	}
}
func CreateMessage(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		chatIDStr := chi.URLParam(r, "id")
		chatID, err := strconv.Atoi(chatIDStr)
		if err != nil{
			http.Error(w, "invalid chat id", http.StatusBadRequest)
			return
		}
		var req struct{
			Text string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		if len(req.Text) == 0 || len(req.Text) > 5000{
			http.Error(w, "text must be 1-5000 chars", http.StatusBadRequest)
			return
		}

		var chat models.Chat
		if err := db.First(&chat, chatID).Error; err != nil{
			if err == gorm.ErrRecordNotFound{
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}
		}

		msg := models.Message{
			ChatID: uint(chatID),
			Text: req.Text,
			CreatedAt: time.Now(),
		}
		if err := db.Create(&msg).Error; err != nil{
			http.Error(w, "failed to create massage", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)

	}
}
func GetChat(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		chatIDStr := chi.URLParam(r, "id")
		chatID, err := strconv.Atoi(chatIDStr)
		if err != nil{
			http.Error(w, "invalid chat id", http.StatusBadRequest)
			return
		} 

		limit := 20 
		if lStr := r.URL.Query().Get("limit"); lStr != ""{
			if l, err := strconv.Atoi(lStr); err == nil && l > 0 && l <= 100 {
                limit = l
            }
        }

        var chat models.Chat
        if err := db.First(&chat, chatID).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                http.Error(w, "chat not found", http.StatusNotFound)
                return
            }
            http.Error(w, "db error", http.StatusInternalServerError)
            return
        }

        var messages []models.Message
        db.Where("chat_id = ?", chatID).Order("created_at desc").Limit(limit).Find(&messages)

        resp := struct {
            Chat     models.Chat      `json:"chat"`
            Messages []models.Message `json:"messages"`
        }{
            Chat: chat,
            Messages: messages,
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
		
	}
}
func DeleteChat(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		chatIDStr := chi.URLParam(r, "id")
        chatID, err := strconv.Atoi(chatIDStr)
        if err != nil {
            http.Error(w, "invalid chat id", http.StatusBadRequest)
            return
        }

        if err := db.Delete(&models.Chat{}, chatID).Error; err != nil {
            http.Error(w, "failed to delete chat", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
	}
}