package model

type Account struct {
	ID             string         `json:"id"`
	AccountNumber  string         `json:"accountNumber"`
	UserID         string         `json:"userID"`
	Balance        float64        `json:"balance"`
	IsMerchant     bool           `json:"isMerchant"`
	MerchantDetail MerchantDetail `json:"merchantDetails,omitempty"`
	IsActive       bool           `json:"isActive"`
}

type MerchantDetail struct {
	BusinessName string `json:"businessName"`
	BusinessType string `json:"businessType"`
}
