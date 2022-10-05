package domain

import (
	"go-banking/banking-app/utils"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRespositoryDB struct {
	client *sqlx.DB
}

func (d AccountRespositoryDB) Save(a Account) (*Account, *utils.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		utils.LogError("Query Error" + err.Error())
		return nil, utils.NewUnexpectedError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		utils.LogError("Query Error" + err.Error())
		return nil, utils.NewUnexpectedError(err.Error())
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}
func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRespositoryDB {
	return AccountRespositoryDB{client: dbClient}
}
