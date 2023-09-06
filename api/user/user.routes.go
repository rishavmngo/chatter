package user

import (
	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/intrf"
	"net/http"
)

func Routes(router *mux.Router, db intrf.Store) {

	user := User{}
	router.HandleFunc("/register", appendDB(db, user.Register)).Methods("GET")
}

type Handler func(http.ResponseWriter, *http.Request, intrf.Store)

func appendDB(db intrf.Store, handler Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db)
	}
}
