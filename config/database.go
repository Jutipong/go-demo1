package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB


func ConnectDB(){

   dsn := "sqlserver://sa:p@ssw0rd@localhost:1433?database=KTBCONRMS_DEV_P"
   db, err :=  gorm.Open(sqlserver.Open(dsn), &gorm.Config{})


   if err != nil {
	   fmt.Println("statuse: ", err)
   }
   //defer DB.Close()
   DB=db

   
}