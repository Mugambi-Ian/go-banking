package domain

import (
	"database/sql"
	"go-banking/banking-auth/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	Register(Login) (*Login, *utils.AppError)
	FindBy(username string, password string) (*Login, *utils.AppError)
	GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *utils.AppError)
	RefreshTokenExists(refreshToken string) *utils.AppError
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (d AuthRepositoryDb) RefreshTokenExists(refreshToken string) *utils.AppError {
	sqlSelect := "select refresh_token from refresh_token_store where refresh_token = ?"
	var token string
	err := d.client.Get(&token, sqlSelect, refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewAuthenticationError("refresh token not registered in the store")
		} else {
			utils.LogError("Unexpected database error: " + err.Error())
			return utils.NewUnexpectedError("unexpected database error")
		}
	}
	return nil
}

func (d AuthRepositoryDb) GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *utils.AppError) {
	var appErr *utils.AppError
	var refreshToken string
	if refreshToken, appErr = authToken.newRefreshToken(); appErr != nil {
		return "", appErr
	}
	sqlInsert := "insert into refresh_token_store (refresh_token) values (?)"
	_, err := d.client.Exec(sqlInsert, refreshToken)
	if err != nil {
		utils.LogError("unexpected database error: " + err.Error())
		return "", utils.NewUnexpectedError("unexpected database error")
	}
	return refreshToken, nil
}

func (d AuthRepositoryDb) Register(l Login) (*Login, *utils.AppError) {
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	sqlInsert := "INSERT INTO users (username, password, role, customer_id, created_on) values (?, ?, ?, ?, ?)"

	_, err := d.client.Exec(sqlInsert, l.Username, l.Password, l.Role, l.CustomerId, createdAt)
	if err != nil {
		utils.LogError("Query Error" + err.Error())
		return nil, utils.NewUnexpectedError(err.Error())
	}

	return &l, nil
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, *utils.AppError) {
	var login Login
	sqlVerify := `SELECT username, password, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM users u
                  LEFT JOIN accounts a ON a.customer_id = u.customer_id
                WHERE username = ?
                GROUP BY a.customer_id`
	err := d.client.Get(&login, sqlVerify, username)

	if err != nil {
		if err == sql.ErrNoRows || utils.ComparePasswords(login.Password, password) {
			return nil, utils.NewAuthenticationError("invalid credentials")
		} else {
			utils.LogError("Error while verifying login request from database: " + err.Error())
			return nil, utils.NewUnexpectedError("Unexpected database error")
		}
	}
	return &login, nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}
