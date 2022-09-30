package domain

import (
	"mugambi-ian/go-banking/errs"
	"mugambi-ian/go-banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRespositoryDB struct {
	client *sqlx.DB
}

func (d AccountRespositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Insert Error:" + err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Insert Error:" + err.Error())
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}
func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRespositoryDB {
	return AccountRespositoryDB{client: dbClient}
}
