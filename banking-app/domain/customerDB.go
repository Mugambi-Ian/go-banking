package domain

import (
	"database/sql"
	"go-banking/banking-app/utils"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (s CustomerRepositoryDB) FindAll(status string) ([]Customer, *utils.AppError) {
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findCustomersSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = s.client.Select(&customers, findCustomersSql)
	} else {
		findCustomersSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = s.client.Select(&customers, findCustomersSql, status)
	}

	if err != nil {
		utils.LogError("Query Error" + err.Error())
		return nil, utils.NewUnexpectedError(err.Error())
	}

	return customers, nil
}

func (s CustomerRepositoryDB) ByID(id string) (*Customer, *utils.AppError) {
	var c Customer
	findCustomersSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	err := s.client.Get(&c, findCustomersSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewNotFoundError("Customer Not Found")
		}

		utils.LogError("Query Error" + err.Error())
		return nil, utils.NewUnexpectedError("Scan Error")
	}

	return &c, nil
}

func (s CustomerRepositoryDB) Create(c Customer) (*Customer, *utils.AppError) {
	sqlInsert := "INSERT INTO customers (name, date_of_birth, city, zipcode, status) values (?, ?, ?, ?, ?)"
	result, err := s.client.Exec(sqlInsert, c.Name, c.DateofBirth, c.City, c.Zipcode, c.Status)
	if err != nil {
		utils.LogError("Query Error" + err.Error())
		return nil, utils.NewUnexpectedError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		utils.LogError("Query Error" + err.Error())
		return nil, utils.NewUnexpectedError(err.Error())
	}
	c.Id = strconv.FormatInt(id, 10)
	return &c, nil
}
func NewCustomerRepositoryDB(dbClient *sqlx.DB) *CustomerRepositoryDB {
	return &CustomerRepositoryDB{client: dbClient}
}
