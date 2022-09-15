package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindBy(username, password string) (*Login, error)
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, error) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM banking.users u
					LEFT JOIN banking.accounts a ON a.customer_id = u.customer_id WHERE username  = ? and password = ? GROUP BY u.customer_id`

	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		} else {
			log.Println("Error while verifying login request from database: ", err.Error())
			return nil, errors.New("unexpected database error")
		}
	}

	return &login, nil
}

func NewAuthRepository(db *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client: db}
}
