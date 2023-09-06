package main

import (
	"github.com/rishavmngo/chatter-backend/api"
	"github.com/rishavmngo/chatter-backend/storage"
)

func main() {

	store := storage.InitilizePostgresStore()
	server := server.Server{}
	server.Initilize(":3000", store)
	server.Run()
}
