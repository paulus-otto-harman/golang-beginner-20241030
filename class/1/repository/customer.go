package repository

import (
	"20241030/class/1/model"
	_interface "20241030/class/1/repository/interface"
	"database/sql"
)

type Customer struct {
	Db *sql.DB
}

func CreateCustomer(db *sql.DB) _interface.CustomerRepository {
	return &Customer{Db: db}
}

func (r *Customer) Create(customer *model.Customer) error {
	query := `INSERT INTO customers (name,user_id) VALUES ($1, $2) RETURNING id`
	err := r.Db.QueryRow(query, customer.Name, customer.User.Id).Scan(&customer.Id)
	if err != nil {
		return err
	}
	return nil
}
