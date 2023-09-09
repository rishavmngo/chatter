package chats

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rishavmngo/chatter-backend/intrf"
	"github.com/rishavmngo/chatter-backend/types"
	"github.com/rishavmngo/chatter-backend/utils"
)

type controller struct {
	db intrf.Store
}

func (controller *controller) GetById(w http.ResponseWriter, r *http.Request) {

	chat_id, _ := strconv.Atoi(mux.Vars(r)["chat_id"])
	chat := types.Chat{}

	chat.ID = uint(chat_id)

	if err := controller.db.GetChatById(&chat); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusOK, chat)
}

func (controller *controller) Add(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	chat := types.Chat{}

	if err := decoder.Decode(&chat); err != nil {
		log.Fatal(err)
	}

	if err := controller.db.AddChat(&chat); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusOK, chat)
}
