package repo

import "database/sql"

type Repository struct {
	Conn *sql.DB
}

type SqlHandler struct {
	err  error
	rows *sql.Rows
}
