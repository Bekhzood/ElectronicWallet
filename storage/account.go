package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type accountRepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) *accountRepo {
	return &accountRepo{db: db}
}

func (a *accountRepo) GetPassword(username string) (string, error) {
	var password string
	query := `SELECT password
		FROM
			users 
		WHERE username = $1`

	rows, err := a.db.Query(query, username)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&password)
	}

	return password, nil
}

func (a *accountRepo) CheckAccount(walletNumber string) (bool, error) {
	var id string
	query := `SELECT id FROM wallets WHERE number = $1`

	err := a.db.QueryRow(query, walletNumber).Scan(&id)

	if err != nil || err == sql.ErrNoRows {
		return false, errors.New("No account found")
	}

	return true, nil
}

func (a *accountRepo) GetHistory() {
	query := `SELECT id
	FROM
		geozones`

	rows, _ := a.db.Query(query)
	var id string
	if rows.Next() {
		rows.Scan(&id)
		fmt.Println(id)
	}
}

func (a *accountRepo) GetBalance(walletNumber string) (int, error) {
	var balance int
	query := `SELECT balance FROM wallets WHERE number = $1`

	err := a.db.QueryRow(query, walletNumber).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (a *accountRepo) UpdateBalance(walletNumber string, additionalMoney int) (rowsAffected int64, err error) {
	var (
		balance    int
		walletType string
	)
	query := `SELECT balance,type FROM wallets WHERE number = $1`
	checkErr := a.db.QueryRow(query, walletNumber).Scan(&balance, &walletType)
	if checkErr != nil {
		return 0, err
	}

	if walletType == "identified" && balance+additionalMoney > 100000 {
		return 0, errors.New("Balance can't be more than 100 000 somoni for identified wallets")
	}

	if walletType == "unidentified" && balance+additionalMoney > 10000 {
		return 0, errors.New("Balance can't be more than 10 000 somoni for unidentified wallets")
	}

	updateQuery := `UPDATE wallets SET balance = balance + $1 WHERE number = $2`

	result, err := a.db.Exec(updateQuery, additionalMoney, walletNumber)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// func (a *accountRepo) CheckAccount(walletNumber string) (bool, error) {
// 	var count int
// 	query := `SELECT COUNT(1) FROM wallets WHERE number = $1`

// 	row, err := a.db.Query(query, walletNumber)

// 	defer row.Close()
// 	if row.Next() {
// 		err = row.Scan(
// 			&count,
// 		)
// 		if err != nil {
// 			return false, err
// 		}
// 	}
// 	if count == 0 {
// 		return false, errors.New("No account found")
// 	}
// 	fmt.Println(count)
// 	return true, nil
// }
