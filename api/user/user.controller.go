package user

import (
	"fmt"
	"net/http"

	"github.com/rishavmngo/chatter-backend/domain"
)

type User struct {
	domain.User
}

func (user User) Register(w http.ResponseWriter, r *http.Request, db domain.Store) {
	err := db.AddUser(user)
	fmt.Printf(err)
	fmt.Fprint(w, "hello from register")
}
