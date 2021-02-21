package main

import (
	 "init/config"
	"init/routers"
	"init/entity"
	
)

var err error

func main() {

	config.ConnectDB()
	config.DB.Table("Customer").AutoMigrate(&entity.Customer{})

	r := routers.SetupRouter()
	r.Run()
}
