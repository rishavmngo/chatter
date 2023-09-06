package user

import (
	"fmt"
	"net/http"

	"github.com/rishavmngo/chatter-backend/domain"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user User) Register(w http.ResponseWriter, r *http.Request, db domain.Store) {
	fmt.Fprint(w, "hello from register")

}
