package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/api/user"
	"github.com/rishavmngo/chatter-backend/domain"
	"net/http"
)

type Server struct {
	DB     domain.Store
	Router *mux.Router
	Port   string
}

func (server *Server) Initilize(port string, store domain.Store) {
	server.DB = store
	server.Router = mux.NewRouter()
	server.Port = port
	server.InitilizeRoutes()
}

func (server *Server) InitilizeRoutes() {

	server.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello,world")

	}).Methods("GET")

	server.Subroute("/user", user.Routes)

}

type InitRouter func(*mux.Router, domain.Store)

func (server *Server) Subroute(path string, initRouter InitRouter) {
	subrouter := server.Router.PathPrefix(path).Subrouter()
	initRouter(subrouter, server.DB)
}

func (server *Server) Run() {
	http.ListenAndServe(server.Port, server.Router)
}
