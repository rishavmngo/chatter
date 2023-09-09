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
	queries = []string{userTable, chatTable, messageTable, recipientsTable, privateChatTable}

	// if _, err := db.Exec(userTable); err != nil {
	// 	log.Fatal(err)
	// }

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

const chatTable = `CREATE TABLE IF NOT EXISTS chats
(
	id SERIAL,
	created_at TIMESTAMP NOT NULL,
	CONSTRAINT chat_pkey PRIMARY KEY (id)
)`

const messageTable = `CREATE TABLE IF NOT EXISTS messages
(
	id SERIAL,
	chat_id INT ,
	author_id INT,
	CONSTRAINT message_pkey PRIMARY KEY (id),
	CONSTRAINT fk_chat
		FOREIGN KEY(chat_id)
			REFERENCES chats(id),
	CONSTRAINT fk_users
		FOREIGN KEY(author_id)
			REFERENCES users(id)
)`

const recipientsTable = `CREATE TABLE IF NOT EXISTS recipients 
(
	id SERIAL,
	message_id INT ,
	receiver_id INT,
	CONSTRAINT recipients_pkey PRIMARY KEY (id),
	CONSTRAINT fk_chat
		FOREIGN KEY(message_id)
			REFERENCES messages(id),
	CONSTRAINT fk_users
		FOREIGN KEY(receiver_id)
			REFERENCES users(id)
)`
const privateChatTable = `CREATE TABLE IF NOT EXISTS private_chat
(
	id SERIAL,
	chat_id INT ,
	member_id INT,
	CONSTRAINT private_chat_pkey PRIMARY KEY (id),
	CONSTRAINT fk_chat
		FOREIGN KEY(chat_id)
			REFERENCES chats(id),
	CONSTRAINT fk_users
		FOREIGN KEY(member_id)
			REFERENCES users(id)
)`

func (postgres *Postgres) AddUser(user *types.User) error {
	err := postgres.db.QueryRow("Insert INTO users(username, email, password) VALUES($1, $2, $3) returning id", user.Username, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (posgres *Postgres) GetUserById() {}
func (postgres *Postgres) GetUserByUsernameAndPassword(user *types.User) error {

	err := postgres.db.QueryRow("SELECT id,email FROM users where username=$1 and password=$2", user.Username, user.Password).Scan(&user.ID, &user.Email)

	if err != nil {
		return err
	}

	return nil
}

func (postgres *Postgres) AddChat(chat *types.Chat) error {

	chat.AddCreatedAt()
	err := postgres.db.QueryRow("INSERT INTO chats(created_at) values($1) returning id", chat.CreatedAt).Scan(&chat.ID)

	if err != nil {
		return err
	}
	return nil
}

func (postgres *Postgres) GetChatById(chat *types.Chat) error {
	err := postgres.db.QueryRow("SELECT created_at FROM chats WHERE id=$1", chat.ID).Scan(&chat.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
