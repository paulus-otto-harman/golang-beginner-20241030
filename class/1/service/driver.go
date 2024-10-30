package service

import (
	"20241030/class/1/model"
	"20241030/class/1/repository"
	"database/sql"
	"errors"
	"fmt"
)

func NewDriver(db *sql.DB, name string, username string, password string) error {
	if name == "" {
		return errors.New("nama tidak boleh kosong")
	}
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	userDB := repository.CreateUser(db)
	user := model.User{
		Username: username,
		Password: password,
	}
	err = userDB.Create(&user)
	if err != nil {
		return err
	}

	driverDB := repository.CreateDriver(db)
	driver := model.Driver{
		Name: name,
		User: user,
	}
	err = driverDB.Create(&driver)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("berhasil input data driver dengan id ", driver.Id)
	return nil
}
