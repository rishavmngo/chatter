package chats

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rishavmngo/chatter-backend/intrf"
	"github.com/rishavmngo/chatter-backend/types"
)

type controller struct {
	db intrf.Store
}

func (controller *controller) GetById(w http.ResponseWriter, r *http.Request) {

}

func (controller *controller) Add(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	chat := types.Chat{}

	if err := decoder.Decode(&chat); err != nil {
		log.Fatal(err)
	}
}
