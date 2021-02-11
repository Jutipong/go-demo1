package service

import (
	"init/entity"
)


type CustomerService interface {
	FindAll() []entity.Customer
}

type customerService struct {
	customers []entity.Customer
}

func New() CustomerService {
	return &customerService{
		customers: []entity.Customer{},
	}
}

func (service *customerService) FindAll() []entity.Customer {
	return service.customers
}

// func (service *videoService) Save(video entity.Video) entity.Video {
// 	service.videos = append(service.videos, video)
// 	return video
// }

// type CustomerService interface {
// 	Find(entity.Customer) []entity.Customer
// }

// type customerService struct {
// 	customers []entity.Customer
// }

// func New() CustomerService {
// 	return &customerService{
// 		customers: []entity.Customer{},
// 	}
// }

// func (service *customerService) Find(result []entity.Customer) {
// 	return
// }
