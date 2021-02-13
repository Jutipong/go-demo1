package controller

import (
	"init/entity"
	"init/helpers"
	"init/service"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

// var validate *validator.Validate

func Find(c *gin.Context) {
	// custom := new(entity.Customer)
	custom := entity.Customer{}
	err := service.Find(&custom)
	if err != nil {
		helpers.RespondJSON(c, 404, custom)
	} else {
		helpers.RespondJSON(c, 200, custom)
	}
}
