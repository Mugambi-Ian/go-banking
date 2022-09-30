package domain

import (
	"database/sql"
	"mugambi-ian/go-banking/errs"
	"mugambi-ian/go-banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (db TransactionRepositoryDB) Transact(transaction Transaction) (*Transaction, *errs.AppError) {
	var a Account
	var balance float64
	findAccountSql := "select customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	err := db.client.Get(&a, findAccountSql, transaction.AccountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account Does not Exist")
		}
		logger.Error("Scan Error: " + err.Error())
		return nil, errs.NewUnexpectedError("Scan Error")
	}
	if transaction.TransactionType == "deposit" {
		balance = a.Amount + transaction.Amount
	} else if transaction.Amount > a.Amount {
		return nil, errs.NewUnexpectedError("Low Balance")
	} else {
		balance = a.Amount - transaction.Amount
	}
	transaction.Balance = balance
	saveTransaction := "INSERT INTO transactions (account_id, Amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, err := db.client.Exec(saveTransaction, transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		logger.Error("Insert Error:" + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Insert Error:" + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	transaction.TransactionId = strconv.FormatInt(id, 10)
	updateAccount := "UPDATE accounts SET amount = ? where account_id = ?"
	_, saveErr := db.client.Exec(updateAccount, balance, transaction.AccountId)
	if err != nil {
		logger.Error("Insert Error:" + saveErr.Error())
		return nil, errs.NewUnexpectedError(saveErr.Error())
	}
	return &transaction, nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepository {
	return TransactionRepositoryDB{client: dbClient}
}
