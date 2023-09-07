package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rishavmngo/chatter-backend/types"
)

type Postgres struct {
	db *sql.DB
}

func InitilizePostgresStore(user, password, dbname, port string) *Postgres {
	postgres := Postgres{}
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)
	var err error
	postgres.db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	ensureTableExist(postgres.db)
	return &postgres
}

var queries []string

func ensureTableExist(db *sql.DB) {
	queries = []string{userTable}

	if _, err := db.Exec(userTable); err != nil {
		log.Fatal(err)
	}

	for _, query := range queries {

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

}

const userTable = `CREATE TABLE IF NOT EXISTS users
(
	id SERIAL,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	email    TEXT,
	CONSTRAINT user_pkey PRIMARY KEY (id),
	CONSTRAINT user_username_unique unique (username),
	CONSTRAINT user_email_unique unique (email)
)`

// handle if the user.Email.String is null
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
func (postgres *Postgres) AddUser(user *types.User) error {
	err := postgres.db.QueryRow("Insert INTO users(username, email, password) VALUES($1, $2, $3) returning id", user.Username, NewNullString(user.Email.String), user.Password).Scan(&user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (posgres *Postgres) GetUserById()       {}
func (posgres *Postgres) GetUserByUsername() {}
