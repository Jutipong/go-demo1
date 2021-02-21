package main

import (
	"init/config"
	"init/entity"
	"init/routers"
	"os"
)

var err error

func main() {

	config.ConnectDB()
	config.DB.Table("Customer").AutoMigrate(&entity.Customer{})

	r := routers.SetupRouter()

	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	r.Run(":" + port)
}
