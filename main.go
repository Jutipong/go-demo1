package main

import (
	 "init/config"
	"init/routers"
	"init/entity"
)

var err error

func main() {

	config.ConnectDB()
	config.DB.AutoMigrate(&entity.Customer{})

	// config.DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/testinger?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	fmt.Println("statuse: ", err)
	// }
	// defer config.DB.Close()

	r := routers.SetupRouter()
	r.Run()
}
