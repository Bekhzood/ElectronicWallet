package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type accountRepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) *accountRepo {
	return &accountRepo{db: db}
}

func (a *accountRepo) GetList() {
	query := `SELECT id
	FROM
		accounts`

	rows, _ := a.db.Query(query)
	var id string
	if rows.Next() {
		rows.Scan(&id)
		fmt.Println(id)
	}
}
