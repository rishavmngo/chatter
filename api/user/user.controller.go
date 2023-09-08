package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rishavmngo/chatter-backend/intrf"
	"github.com/rishavmngo/chatter-backend/jwtUtils"
	"github.com/rishavmngo/chatter-backend/types"
	"github.com/rishavmngo/chatter-backend/utils"
)

type Controller struct {
	db intrf.Store
}

func (controller *Controller) Register(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	user := types.User{}
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	err = controller.db.AddUser(&user)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, user)
}

func (controller *Controller) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	user := types.User{}
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	err = controller.db.GetUserByUsernameAndPassword(&user)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	token := jwtUtils.GenerateToken(user.ID)
	utils.RespondWithJson(w, http.StatusOK, map[string]string{"token": token})
}
