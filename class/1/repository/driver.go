package repository

import (
	"20241030/class/1/model"
	_interface "20241030/class/1/repository/interface"
	"database/sql"
)

type Driver struct {
	Db *sql.DB
}

func CreateDriver(db *sql.DB) _interface.DriverRepository {
	return &Driver{Db: db}
}

func (r *Driver) Create(driver *model.Driver) error {
	query := `INSERT INTO drivers (name,user_id) VALUES ($1, $2) RETURNING id`
	err := r.Db.QueryRow(query, driver.Name, driver.User.Id).Scan(&driver.Id)
	if err != nil {
		return err
	}
	return nil
}
