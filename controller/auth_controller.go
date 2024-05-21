package controller

import (
	"basecode/services"
	"basecode/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (controller *AuthController) Register(c *gin.Context) {
	request := new(web.RegisterUserRequest)

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	status, err := controller.AuthService.Register(request)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{"message": "success"})
}

func (controller *AuthController) Login(c *gin.Context) {
	request := new(web.LoginUserRequest)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response, status, err := controller.AuthService.Login(*request)
	if err != nil {
		c.JSON(status, gin.H{"message": err.Error()})
	}

	c.JSON(status, gin.H{"data": response})
}
