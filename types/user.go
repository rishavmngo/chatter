package types

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
