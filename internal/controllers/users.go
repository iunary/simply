package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iunary/simply/internal/services"
	"github.com/spf13/viper"
)

type UserController struct {
	logger *log.Logger
	v      *viper.Viper
	svc    services.IUserService
}

func NewUserController(logger *log.Logger, v *viper.Viper, svc services.IUserService) *UserController {
	return &UserController{
		logger: logger,
		svc:    svc,
		v:      v,
	}
}

func (uc *UserController) Get(c *gin.Context) {
	uc.logger.Println("User controller get user")
	users, err := uc.svc.All(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not found",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
