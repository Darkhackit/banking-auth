package domain

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"

type Claims struct {
	CustomerID string `json:"customer_id"`
	AccountID  string `json:"account_id"`
	Username   string `json:"username"`
	Expiry     string `json:"expiry"`
	Role       string `json:"role"`
}

func (c Claims) IsUserRole() bool {
	if c.Role == "user" {
		return true
	}
	return false
}
