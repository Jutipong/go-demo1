package service

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func Login(username string, password string) bool {
	return true
	// return service.authorizedUsername == username &&
	// 	service.authorizedPassword == password
}
