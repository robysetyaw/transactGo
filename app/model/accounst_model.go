package model

type Account struct {
	AccountNumber  string         `json:"accountNumber"`
	UserID         string         `json:"userID"`
	Balance        float64        `json:"balance"`
	IsMerchant     bool           `json:"isMerchant"`
	MerchantDetail MerchantDetail `json:"merchantDetails,omitempty"`
}

type MerchantDetail struct {
	BusinessName string `json:"businessName"`
	BusinessType string `json:"businessType"`
}