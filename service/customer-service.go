package service

import (
	"init/entity"
)

func FindALL() []entity.Customer {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func FindALL() (err error) {
	if err = config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}
