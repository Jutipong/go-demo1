package controller

import (
	"github.com/gin-gonic/gin"
	"init/model"
	"init/service"
)

// type LoginController interface {
// 	Login(ctx *gin.Context) string
// }

// type loginController struct {
// 	loginService service.LoginService
// 	jWtService   service.JWTService
// }

// func NewLoginController(loginService service.LoginService,
// 	jWtService service.JWTService) LoginController {
// 	return &loginController{
// 		loginService: loginService,
// 		jWtService:   jWtService,
// 	}
// }

func Login(ctx *gin.Context) string {
	var credentials model.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := service.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return service.GenerateToken(credentials.Username, true)
	}
	return ""
}
