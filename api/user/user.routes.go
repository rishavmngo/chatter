package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/intrf"
)

func Routes(router *mux.Router, db intrf.Store) {

	router.HandleFunc("/register", appendDB(db, Register)).Methods("POST")
}

type Handler func(http.ResponseWriter, *http.Request, intrf.Store)

func appendDB(db intrf.Store, handler Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db)
	}
}
