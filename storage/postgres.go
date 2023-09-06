package storage

import "database/sql"

type Postgres struct {
	db *sql.DB
}

func InitilizePostgresStore() *Postgres {
	return &Postgres{nil}
}
