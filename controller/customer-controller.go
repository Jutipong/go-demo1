package controller

import (
	"init/entity"
	"init/helpers"
	"init/service"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

// var validate *validator.Validate

func FindID(c *gin.Context) {
	// custom := new(entity.Customer)
	id :=c.Params.ByName("id")
	custom := entity.Customer{}
	err := service.FindID(&custom, id)
	if err != nil {
		helpers.RespondJSON(c, 404, err.Error(), custom)
	} else {
		helpers.RespondJSON(c, 200, "success", custom)
	}

}


func FindAll(c *gin.Context) {
	// custom := new(entity.Customer)
	custom := []entity.Customer{}
	err := service.FindAll(&custom)
	if err != nil {
		helpers.RespondJSON(c, 404, err.Error(), custom)
	} else {
		helpers.RespondJSON(c, 200, "success", custom)
	}
}


func AddNewCustomer(c *gin.Context) {
	var custom entity.Customer
	err := c.ShouldBindJSON(&custom)
	c.BindJSON(&custom)

	if err != nil {
		helpers.RespondJSON(c, 404, err.Error(), custom)
	} else{
		err = service.AddNewCustomer(&custom)
		if err != nil {
			helpers.RespondJSON(c, 404, err.Error() ,custom)
		} else {
			helpers.RespondJSON(c, 200, "success" ,custom)
		}
	}	
}

func PutOneCustomer(c *gin.Context) {
	var custom entity.Customer
	id := c.Params.ByName("id")
	err := service.FindID(&custom, id)
	if err != nil {
		helpers.RespondJSON(c, 404,err.Error(), custom)
	}
	c.BindJSON(&custom)
	err = service.PutOneCustomer(&custom, id)
	if err != nil {
		helpers.RespondJSON(c, 404,err.Error(), custom)
	} else {
		helpers.RespondJSON(c, 200, "success", custom)
	}
}

func DeleteCustomer(c *gin.Context) {
	var custom entity.Customer
	id := c.Params.ByName("id")
	err := service.DeleteCustomer(&custom, id)
	if err != nil {
		helpers.RespondJSON(c, 404,err.Error(), custom)
	} else {
		helpers.RespondJSON(c, 200, "success", custom)
	}
}