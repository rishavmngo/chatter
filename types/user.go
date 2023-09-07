package types

import (
	"net/http"
)

type User struct {
	ID       uint           `json:"id"`
	Username string         `json:"username"`
	Email    NullableString `json:"email"`
	Password string         `json:"password"`
}

type UserController interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}
