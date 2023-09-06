package domain

type Store interface {
	GetUserById()
	GetUserByUsername()
	AddUser()
}
