package model

type Order struct {
	Id       int
	Customer Customer
	Area     Area
	Fare     int
	Driver   Driver
}
