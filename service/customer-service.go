package service

import (
	"init/entity"
)

type CustomerService interface {
	FindAll() []entity.Customer
}

type customerService struct {
	customers []entity.Customer
}

func New() CustomerService {
	return &customerService{
		customers: []entity.Customer{},
	}
}

func (service *customerService) FindAll() []entity.Customer {
	var result = []entity.Customer{
		{
			Name: "Alice",
			Age:  12,
		},
		{
			Name: "Bob",
			Age:  15,
		},
	}
	return result
}
