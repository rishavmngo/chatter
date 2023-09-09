package types

import (
	"net/http"
	"time"
)

type Chat struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatController interface {
	GetById(http.ResponseWriter, *http.Request)
	Add(http.ResponseWriter, *http.Request)
}

func (chat *Chat) AddCreatedAt() {
	chat.CreatedAt = time.Now()
}
