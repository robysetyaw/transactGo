package model

type Payment struct {
	ID         int    `json:"id"`
	UserID     int    `json:"userId"`
	MerchantID int    `json:"merchantId"`
	Amount     int    `json:"amount"`
	Timestamp  string `json:"timestamp"`
}
