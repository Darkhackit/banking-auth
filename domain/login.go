package domain

import (
	"database/sql"
	"errors"
	"github.com/Darkhackit/banking-auth/errs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const TokenDuration = time.Hour

type Login struct {
	ID         string         `db:"user_id"`
	CustomerID sql.NullString `db:"customer_id"`
	Username   string         `db:"username"`
	Password   string         `db:"password"`
	Role       string         `db:"role"`
	AccountID  sql.NullString `db:"account_id"`
}

func (l Login) GenerateToken() (*string, error) {
	var claims jwt.MapClaims
	if l.AccountID.Valid && l.CustomerID.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		return nil, errors.New("cannot generate token")
	}
	return &signedTokenAsString, nil
}
func (l Login) claimsForUser() jwt.MapClaims {
	//accounts := strings.Split(l.AccountID.String, ',')
	return jwt.MapClaims{
		"customer_id": l.CustomerID.String,
		"username":    l.Username,
		"role":        l.Role,
		"account_id":  l.AccountID.String,
		"exp":         time.Now().Add(TokenDuration).Unix(),
	}
}
func (l Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"role": l.Role,
		"user": l.Username,
		"exp":  time.Now().Add(TokenDuration).Unix(),
	}
}

type LoginRepository interface {
	Login(login Login) (*Login, *errs.AppError)
	Register(login Login) (*Login, *errs.AppError)
}
