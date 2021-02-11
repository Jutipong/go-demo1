package service

import (
	"fmt"
	"init/entity"
)

func FindByID(customer entity.Customer) entity.Customer {
	fmt.Println("customer service: ", customer)
	return customer
}
