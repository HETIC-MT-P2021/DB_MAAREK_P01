package repo

import "database/sql"

type Repository struct {
	Conn *sql.DB
}