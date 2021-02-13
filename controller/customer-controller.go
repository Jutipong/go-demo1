package controller

import (
	"init/helpers"
	"init/service"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

// var validate *validator.Validate

func FindALL(c *gin.Context) {
	 := service.FindALL()
	if err != nil {
		helpers.RespondJSON(c, 404, book)
	} else {
		helpers.RespondJSON(c, 200, book)
	}
}