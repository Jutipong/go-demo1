package service

import (
	"init/entity"
	"init/config"
)

func FindAll(c *[]entity.Customer) (err error) {
	if err = config.DB.Table("Customer").Find(&c).Error; err != nil {
		return err
	}
	return nil
}

func FindID(c *entity.Customer, id string) (err error) {
	if err = config.DB.Table("Customer").Where("id = ?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}


func AddNewCustomer(c *entity.Customer) (err error) {
	if err = config.DB.Table("Customer").Create(c).Error; err != nil {
		return err
	}
	return nil
}


func PutOneCustomer(c *entity.Customer, id string) (err error) {
	
	config.DB.Table("Customer").Where("id = ?", id).Save(c)
	return nil
}

func DeleteCustomer(c *entity.Customer, id string) (err error) {
	config.DB.Table("Customer").Where("id = ?", id).Delete(c)
	return nil
}