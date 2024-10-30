package repository_interface

import "20241030/class/1/model"

type OrderRepository interface {
	Get1(query string) (interface{}, error)
	Get2(query string) (interface{}, error)
	Get3(query string) (interface{}, error)
	Get4(query string) (interface{}, error)
	Get5(query string) (int, error)
	Create(order *model.Order) error
	//Update(customer *model.Customer) error
	//Delete(id int) error
	//GetByID(id int) (*model.Customer, error)
	//GetAll() ([]*model.Customer, error)
}
