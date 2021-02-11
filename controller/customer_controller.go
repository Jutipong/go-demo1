package controller

import (
	"fmt"
	"init/entity"
	"init/service"
)

func FindALL(cust entity.Customer) (result entity.Customer) {
	fmt.Println("controller cust:", cust)
	result = service.FindByID(cust)
	return result
}
