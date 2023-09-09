package intrf

import (
	"github.com/rishavmngo/chatter-backend/types"
)

type Store interface {
	GetUserById()
	GetUserByUsernameAndPassword(*types.User) error
	AddUser(*types.User) error
	AddChat(*types.Chat) error
	GetChatById(*types.Chat) error
}
