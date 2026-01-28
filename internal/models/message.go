package models

import "time"


type Message struct{
	ID uint `gorm:"primaryKey"`
	ChatID uint `gorm:"not null;index"`
	Text string `gorm:"size:5000;not null"`
	CreatedAt time.Time
}