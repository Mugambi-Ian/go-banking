package domain

import (
	"go-banking/banking-app/utils"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (db TransactionRepositoryDB) Transact(t Transaction) (*Transaction, *utils.AppError) {
	tx, err := db.client.Begin()
	if err != nil {
		utils.LogError("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, utils.NewUnexpectedError("Unexpected database error")
	}

	result, err := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if err != nil {
		tx.Rollback()
		utils.LogError("Error while saving transaction: " + err.Error())
		return nil, utils.NewUnexpectedError("Unexpected database error")
	}

	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		utils.LogError("Error while saving transaction: " + err.Error())
		return nil, utils.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		utils.LogError("Error while commiting transaction for bank account: " + err.Error())
		return nil, utils.NewUnexpectedError("Unexpected database error")
	}
	transactionId, transactError := result.LastInsertId()

	if transactError != nil {
		utils.LogError("Error while getting the last transaction id: " + transactError.Error())
		return nil, utils.NewUnexpectedError("Unexpected database error")
	}

	var a Account
	var c Customer
	findAccountSql := "select customer_id, account_type, amount from accounts where account_id = ?"
	findCustomerSql := "select customer_id, name from customers where customer_id = ?"
	err = db.client.Get(&a, findAccountSql, t.AccountId)
	err = db.client.Get(&c, findCustomerSql, a.CustomerId)

	if err != nil {
		utils.LogError("Error while getting the user account: " + err.Error())
		return nil, utils.NewUnexpectedError("Unexpected database error")
	}
	t.Account = a
	t.Customer = c
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	return &t, nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepository {
	return TransactionRepositoryDB{client: dbClient}
}
