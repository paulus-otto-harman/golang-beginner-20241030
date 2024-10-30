package repository

import (
	"20241030/class/1/model"
	_interface "20241030/class/1/repository/interface"
	"database/sql"
)

type User struct {
	Db *sql.DB
}

func CreateUser(db *sql.DB) _interface.UserRepository {
	return &User{Db: db}
}

func (r *User) Create(user *model.User) error {
	query := `INSERT INTO users (username,password) VALUES ($1, $2) RETURNING id`
	err := r.Db.QueryRow(query, user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}
