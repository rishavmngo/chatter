package user

import (
	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/domain"
	"net/http"
)

func Routes(router *mux.Router, db domain.Store) {

	user := User{}
	router.HandleFunc("/register", appendDB(db, user.Register)).Methods("GET")
}

type Handler func(http.ResponseWriter, *http.Request, domain.Store)

func appendDB(db domain.Store, handler Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db)
	}
}
