package user

import (
	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/intrf"
	"github.com/rishavmngo/chatter-backend/types"
)

func Routes(router *mux.Router, db intrf.Store) {
	var user types.UserController

	user = &Controller{db}

	router.HandleFunc("/register", user.Register).Methods("POST")
	router.HandleFunc("/login", user.Login).Methods("POST")
}
