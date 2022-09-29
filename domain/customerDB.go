package domain

import (
	"database/sql"
	"log"
	"mugambi-ian/go-banking/errs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (s CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error
	if status == "" {
		findCustomersSql := "select  customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = s.client.Query(findCustomersSql)
	} else {
		statusValue := 0
		if status == "active" {
			statusValue = 1
		}
		findCustomersSql := "select  customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = s.client.Query(findCustomersSql, statusValue)
	}

	if err != nil {
		log.Println("Query Error" + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Scan Error" + err.Error())
			return nil, errs.NewUnexpectedError(err.Error())
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (s CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	findCustomersSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := s.client.QueryRow(findCustomersSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewUnexpectedError("Customer Not Found")
		}
		log.Println("Scan Error" + err.Error())
		return nil, errs.NewNotFoundError("Scan Error")
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}
