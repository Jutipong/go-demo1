package main

import (
	"fmt"
	"init/controller"
	"init/entity"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/customer", func(c *gin.Context) {
		cust :=  entity.Customer{
			ID:  "12345",
			Code: "C0001",
			Name: "Jutipong",
			Age: 28,
			Email: "xxxxxx.12345@gmail.com",
		} 
		fmt.Println("initial obj customer: ", cust)
		result := controller.FindALL(cust)
		c.JSONP(200,result) 
	})
	r.Run()
}
