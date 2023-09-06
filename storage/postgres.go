package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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

func (postgres *Postgres) AddUser() {

}

func (posgres *Postgres) GetUserById()       {}
func (posgres *Postgres) GetUserByUsername() {}
