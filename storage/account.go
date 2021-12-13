package storage

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Bekhzood/ElectronicWallet/models"
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
	var number string
	query := `SELECT id FROM wallets WHERE number = $1`

	err := a.db.QueryRow(query, walletNumber).Scan(&number)

	if err != nil || err == sql.ErrNoRows {
		return false, errors.New("No account found")
	}

	return true, nil
}

func (a *accountRepo) GetHistory(number string) (res models.WalletTransactionsHistory, err error) {
	var (
		transaction models.Transaction
		count       int
	)
	now := time.Now()
	today := now.Format("2006-01-02 15:04:05.000000")
	lastMonth := now.AddDate(0, -1, 0).Format("2006-01-02 15:04:05.000000")

	query := `SELECT *, COUNT(*) OVER () AS total_count FROM transactions_history`
	filter := " WHERE (created_at >= $1) AND (created_at <= $2) AND (sender_wallet_number = $3 OR receiver_wallet_number = $3)"
	order := " ORDER BY created_at"
	arrangement := " DESC"

	completeQuery := query + filter + order + arrangement

	rows, err := a.db.Query(completeQuery, lastMonth, today, number)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&transaction.ID,
			&transaction.Sum,
			&transaction.SenderWalletNumber,
			&transaction.ReceiverWalletNumber,
			&transaction.CreatedAt,
			&count,
		)
		if err != nil {
			return res, err
		}
		res.Transactions = append(res.Transactions, transaction)
	}
	res.Count = count
	return res, nil
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

func (a *accountRepo) UpdateBalance(senderWallet int, updateInfo models.UpdateBalance) (rowsAffected int64, err error) {
	var (
		senderBalance   int
		receiverBalance int
		walletType      string
	)
	// checking sender's wallet number is not equal to receiver's one
	if senderWallet == updateInfo.ReceiverWalletNumber {
		return 0, errors.New("Cannot send to its own wallet")
	}

	// getting sender's balance
	senderQuery := `SELECT balance FROM wallets WHERE number = $1`
	err = a.db.QueryRow(senderQuery, senderWallet).Scan(&senderBalance)
	if err != nil {
		return 0, err
	}

	// checking sender's balance to be not negative after transaction
	if senderBalance-updateInfo.Sum < 0 {
		return 0, errors.New("Not enough money")
	}

	// getting receiver's balance
	receiverQuery := `SELECT balance,type FROM wallets WHERE number = $1`
	err = a.db.QueryRow(receiverQuery, updateInfo.ReceiverWalletNumber).Scan(&receiverBalance, &walletType)
	if err != nil {
		return 0, err
	}

	// checking sum to be not negative
	if updateInfo.Sum < 0 {
		return 0, errors.New("Sending sum cannot be negative")
	}

	// checking max sum for identified wallet
	if walletType == "identified" && receiverBalance+updateInfo.Sum > 100000 {
		return 0, errors.New("Balance can't be more than 100 000 somoni for identified wallets")
	}

	// checking max sum for unidentified wallet
	if walletType == "unidentified" && receiverBalance+updateInfo.Sum > 10000 {
		return 0, errors.New("Balance can't be more than 10 000 somoni for unidentified wallets")
	}

	updateReceiverQuery := `UPDATE wallets SET balance = balance + $1 WHERE number = $2`

	result1, err := a.db.Exec(updateReceiverQuery, updateInfo.Sum, updateInfo.ReceiverWalletNumber)
	if err != nil {
		return 0, err
	}

	rowsAffected1, err := result1.RowsAffected()
	if err != nil {
		return 0, err
	}

	updateSenderQuery := `UPDATE wallets SET balance = balance - $1 WHERE number = $2`

	result2, err := a.db.Exec(updateSenderQuery, updateInfo.Sum, senderWallet)
	if err != nil {
		return 0, err
	}

	rowsAffected2, err := result2.RowsAffected()
	if err != nil {
		return 0, err
	}

	transactionsQuery := `INSERT INTO transactions_history(id, sum,sender_wallet_number,receiver_wallet_number) VALUES($1,$2,$3,$4)`

	_, err = a.db.Exec(transactionsQuery, updateInfo.ID, updateInfo.Sum, senderWallet, updateInfo.ReceiverWalletNumber)
	if err != nil {
		return 0, err
	}

	rowsAffected = rowsAffected1 + rowsAffected2
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
