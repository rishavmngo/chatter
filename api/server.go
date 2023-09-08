package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/api/chats"
	"github.com/rishavmngo/chatter-backend/api/user"
	"github.com/rishavmngo/chatter-backend/intrf"
)

type Server struct {
	DB     intrf.Store
	Router *mux.Router
	Port   string
}

func (server *Server) Initilize(port string, store intrf.Store) {
	server.DB = store
	server.Router = mux.NewRouter()
	server.Port = port
	server.InitilizeRoutes()
}

func (server *Server) InitilizeRoutes() {

	server.Subroute("/user", user.Routes)
	server.Subroute("/chats", chats.Routes)
}

type InitRouterType func(*mux.Router, intrf.Store)

func (server *Server) Subroute(path string, initRouter InitRouterType) {
	subrouter := server.Router.PathPrefix(path).Subrouter()
	initRouter(subrouter, server.DB)
}

func (server *Server) Run() {
	http.ListenAndServe(server.Port, server.Router)
}
