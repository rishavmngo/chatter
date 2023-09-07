package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rishavmngo/chatter-backend/intrf"
	"github.com/rishavmngo/chatter-backend/types"
	"github.com/rishavmngo/chatter-backend/utils"
)

func Register(w http.ResponseWriter, r *http.Request, db intrf.Store) {
	decoder := json.NewDecoder(r.Body)

	user := types.User{}
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AddUser(&user)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, user)
}

func Login(w http.ResponseWriter, r *http.Request, db intrf.Store) {

}
