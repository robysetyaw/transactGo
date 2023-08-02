package model

import "time"

type Transaction struct {
	ID               string    `json:"id"`
	FromAccountNumber string    `json:"fromAccountNumber"`
	ToAccountNumber   string    `json:"toAccountNumber"`
	Amount            float64   `json:"amount"`
	Timestamp         time.Time `json:"timestamp"`
	Description       string    `json:"description"`
	TxType            string    `json:"txType"` // Transaction type: "transfer" or "payment"
}
