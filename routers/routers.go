package routers

import (
	// "init/controller"
	"net/http"

	"init/controller"
	"init/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// v1 := r.Group("/v1")
	// {
	// 	// v1.GET("book", controller.Find)
	// 	// v1.POST("book", Controllers.AddNewBook)
	// 	// v1.GET("book/:id", Controllers.GetOneBook)
	// 	// v1.PUT("book/:id", Controllers.PutOneBook)
	// 	// v1.DELETE("book/:id", Controllers.DeleteBook)
	// }

	r.POST("/login", func(ctx *gin.Context) {
		token := controller.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := r.Group("/api", middlewares.AuthorizeJWT())
	// apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/customer", controller.FindAll)
		// {
		// 	// result := _customerController.FindAll()
		// 	result :=
		// 	c.JSON(200, result)
		// })
		apiRoutes.GET("/customer/:id", controller.FindID)
		apiRoutes.POST("customer", controller.AddNewCustomer)
		apiRoutes.PUT("customer/:id", controller.PutOneCustomer)
		apiRoutes.DELETE("customer/:id", controller.DeleteCustomer)
	}
	return r
}
