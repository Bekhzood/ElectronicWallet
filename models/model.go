package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"Username"`
	Password  int       `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Wallet struct {
	ID        string    `json:"id"`
	Number    int       `json:"number"`
	Balance   int       `json:"balance"`
	UserId    string    `json:"user_id"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	ID               int       `json:"id"`
	Sum              int       `json:"sum"`
	SenderWalletId   string    `json:"sender_wallet_id"`
	ReceiverWalletId string    `json:"receiver_wallet_id"`
	CreatedAt        time.Time `json:"created_at"`
}

type WalletTransactionsHistory struct {
	Count        int           `json:"count"`
	Transactions []Transaction `json:"transactions"`
}
