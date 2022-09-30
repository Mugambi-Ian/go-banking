package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"mugambi-ian/go-banking/errs"
	"mugambi-ian/go-banking/logger"
	"os"
	"time"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (s CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
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
		logger.Error("Query Error" + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return customers, nil
}

func (s CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	var c Customer
	findCustomersSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	err := s.client.Get(&c, findCustomersSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		}
		logger.Error("Scan Error" + err.Error())
		return nil, errs.NewUnexpectedError("Scan Error")
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword != "" {
		dbUser += ":" + dbPassword
	}

	client, err := sqlx.Open("mysql", dbUser+"@tcp("+dbAddress+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}
