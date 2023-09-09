package messages

import (
	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/intrf"
)

func Routes(router *mux.Router, db intrf.Store) {

	messages := Messages{db}
	router.HandleFunc("/add", messages.Add)
	router.HandleFunc("/getById/{message_id}", messages.Add)

}
