package controller

import (
	"init/entity"
	"init/service"

	"github.com/go-playground/validator/v10"
)

type CustomerController interface {
	FindAll() []entity.Customer
}

type controller struct {
	service service.CustomerService
}

var validate *validator.Validate

func New(service service.CustomerService) CustomerController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Customer {
	return c.service.FindAll()
}

// type CustomerController interface {
// 	Find(ctx *gin.Context) []entity.Customer
// 	// Insert(ctx *gin.Context) bool
// 	// Update(ctx *gin.Context) bool
// 	// Delete(ctx *gin.Context) bool
// }

// type controller struct {
// 	service service.CustomerService
// }

// //##register validate
// var validate *validator.Validate
// func New(service service.CustomerService) CustomerController {
// 	validate = validator.New()
// 	return &controller{
// 		service: service,
// 	}
// }

// func (c *controller) Find(ctx *gin.Context) model.Reponse {
// 	cust := new(entity.Customer) //entity.Customer

// 	fmt.Println("cust1 : ", cust)

// 	fmt.Println("cust1 : ", cust)
// 	// var result = new(model.Reponse)

// 	//##check  map data json
// 	err := ctx.ShouldBindJSON(&cust)

// 	fmt.Println("err : ", err)
// 	fmt.Println("cust1 : ", cust)

// 	fmt.Println("cust2 : ", cust)

// 	if err != nil {
// 		return model.Reponse{
// 			Message: "Error",
// 		}
// 	}

// 	return model.Reponse{
// 		Status: true ,
// 		Code: 200 ,
// 		Datas: cust,
// 	}

// 	//##validate model
// 	// err = validate.Struct(cust)
// 	// if err != nil {
// 	// 	return result
// 	// }
// 	// res := c.service.Find(cust);

// }
