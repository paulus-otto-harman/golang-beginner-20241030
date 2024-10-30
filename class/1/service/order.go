package service

import (
	"20241030/class/1/model"
	"20241030/class/1/repository"
	"database/sql"
	"fmt"
)

func NewOrder(db *sql.DB, customer model.Customer, area model.Area, fare int, driver model.Driver) error {
	orderDB := repository.CreateOrder(db)
	order := model.Order{
		Customer: customer,
		Area:     area,
		Fare:     fare,
		Driver:   driver,
	}
	err = orderDB.Create(&order)
	if err != nil {
		return err
	}
	return nil
}

func Get1(db *sql.DB, query string) {
	orderDB := repository.Get(db)
	orderDB.Get1(query)
	fmt.Println(orderDB.Get1(query))
}

func Get2(db *sql.DB, query string) {
	orderDB := repository.Get(db)
	orderDB.Get2(query)
	fmt.Println(orderDB.Get2(query))
}

func Get3(db *sql.DB, query string) {
	orderDB := repository.Get(db)
	orderDB.Get3(query)
	fmt.Println(orderDB.Get3(query))
}

func Get4(db *sql.DB, query string) {
	orderDB := repository.Get(db)
	orderDB.Get4(query)
	fmt.Println(orderDB.Get2(query))
}

func Get5(db *sql.DB, query string) {
	orderDB := repository.Get(db)
	orderDB.Get5(query)
	fmt.Println(orderDB.Get2(query))
}
