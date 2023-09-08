package chats

import (
	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/intrf"
	"github.com/rishavmngo/chatter-backend/types"
)

func Routes(router *mux.Router, db intrf.Store) {
	var chat types.ChatController

	chat = &controller{db}

	router.HandleFunc("/add", chat.Add).Methods("POST")
	router.HandleFunc("/getbyid/{chat_id}", chat.GetById).Methods("GET")
}
