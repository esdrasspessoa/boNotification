package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Message string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewNotification(title, message string) *Notification {
	return &Notification{
		ID: uuid.New(),
		Title: title,
		Message: message,
		CreatedAt: time.Now(),
	}
}
