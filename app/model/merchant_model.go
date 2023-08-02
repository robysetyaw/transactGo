package model

type Merchant struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	UserID       string `json:"userId"`
	BusinessName string `json:"businessName"`
	BusinessType string `json:"businessType"`
}
