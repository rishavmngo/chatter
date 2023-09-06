package intrf

import (
	"github.com/rishavmngo/chatter-backend/types"
)

type Store interface {
	GetUserById()
	GetUserByUsername()
	AddUser(*types.User) error
}
