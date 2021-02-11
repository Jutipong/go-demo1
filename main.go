package main

import (
	"init/controller"
	"init/service"

	"github.com/gin-gonic/gin"
)

var (
	_customerService    service.CustomerService       = service.New()
	_customerController controller.CustomerController = controller.New(_customerService)

	// loginService service.LoginService = service.NewLoginService()
	// jwtService   service.JWTService   = service.NewJWTService()

	// loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {
	r := gin.Default()

	r.GET("/customer", func(c *gin.Context) {
		result := _customerController.FindAll()
		c.JSON(200, result)
	})

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// port := os.Getenv("PORT")
	// Elastic Beanstalk forwards requests to port 5000
	// if port != "" {
	// 	port =
	// }
	r.Run()
}
