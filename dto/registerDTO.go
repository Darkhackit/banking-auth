package dto

import "database/sql"

type RegisterDTO struct {
	CustomerID sql.NullString `json:"customer_id"`
	AccountID  sql.NullString `json:"account_id"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Role       string         `json:"role"`
}
