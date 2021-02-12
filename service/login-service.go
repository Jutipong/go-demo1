package service

import "fmt"

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "admin",
		authorizedPassword: "1234",
	}
}

func (service *loginService) Login(username string, password string) bool {
	u := service.authorizedUsername
	p := service.authorizedPassword
	fmt.Println(u)
	fmt.Println(p)

	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
