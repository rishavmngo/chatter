package types

import "net/http"

type Message struct {
	ID       uint `json:"id"`
	ChatID   uint `json:"chat_id"`
	AuthorID uint `json:"author_id"`
}

type MessagesController interface {
	Add(http.ResponseWriter, *http.Request)
	GetById(http.ResponseWriter, *http.Request)
}
