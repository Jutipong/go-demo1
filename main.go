package main

import (
	"init/controller"
	"init/middlewares"
	"init/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	_customerService    service.CustomerService       = service.New()
	_customerController controller.CustomerController = controller.New(_customerService)

	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := r.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/customer", func(c *gin.Context) {
			result := _customerController.FindAll()
			c.JSON(200, result)
		})
	}

	r.Run()
}
