package messages

import (
	"net/http"

	"github.com/rishavmngo/chatter-backend/intrf"
)

type Messages struct {
	db intrf.Store
}

func (messages *Messages) Add(w http.ResponseWriter, r *http.Request) {

}

func (messages *Messages) GetById(w http.ResponseWriter, r *http.Request) {}
