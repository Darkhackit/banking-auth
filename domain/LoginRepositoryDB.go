package domain

import (
	"database/sql"
	"errors"
	"github.com/Darkhackit/banking-auth/errs"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type LoginRepositoryDB struct {
	client *sqlx.DB
}

func (lr LoginRepositoryDB) Register(login Login) (*Login, *errs.AppError) {
	tx, err := lr.client.Begin()
	if err != nil {
		return nil, errs.NewUnexpectedError("error beginning transaction" + err.Error())
	}
	registerSql := `INSERT INTO users (username, customer_id, password,role,account_id) VALUES (?, ?, ?,?,?)`
	result, err := tx.Exec(registerSql, login.Username, login.CustomerID, login.Password, login.Role, login.AccountID)
	if err != nil {
		_ = tx.Rollback()
		return nil, errs.NewUnexpectedError("error inserting user " + err.Error())
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
	}
	login.ID = strconv.FormatInt(lastInsertId, 10)
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
	}
	return &login, nil
}

func (lr LoginRepositoryDB) Login(login Login) (*Login, *errs.AppError) {
	var user Login
	loginSql := `SELECT * FROM users WHERE username = ?`
	err := lr.client.Get(&user, loginSql, login.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFoundError("user not found")
		}
		return nil, errs.NewUnexpectedError("error getting user" + err.Error())
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte([]byte(login.Password))); err != nil {
		return nil, errs.NewValidationError("wrong password")
	}
	return &user, nil
}

func NewLoginRepositoryDB(dbClient *sqlx.DB) LoginRepositoryDB {
	return LoginRepositoryDB{dbClient}
}
