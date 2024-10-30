package repository_interface

import "20241030/class/1/model"

type CustomerRepository interface {
	Create(customer *model.Customer) error
	//Update(customer *model.Customer) error
	//Delete(id int) error
	//GetByID(id int) (*model.Customer, error)
	//GetAll() ([]*model.Customer, error)
}
