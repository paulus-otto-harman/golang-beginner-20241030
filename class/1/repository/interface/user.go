package repository_interface

import "20241030/class/1/model"

type UserRepository interface {
	Create(user *model.User) error
	//Update(user *model.User) error
	//Delete(id int) error
	//GetByID(id int) (*model.User, error)
	//GetAll() ([]*model.User, error)
}
