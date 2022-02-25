package controller

import (
	"fmt"
	"net/http"
	"service/internal/config"
	"service/internal/model"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Config 			*config.Configuration
	UserService *service.UserService
}

func NewUserController(
	cfg *config.Configuration,
	us  *service.UserService,
) *UserController {
	return &UserController{
		Config:      cfg,
		UserService: us,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var reqUser model.Users
	if err := ctx.ShouldBindJSON(&reqUser); err != nil {
		fmt.Println("User controller create ShouldBindJSON error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("reqUser : ", reqUser)

	fmt.Println("userService : ", c.UserService)

	if err := c.UserService.CreateUser(&reqUser); err != nil {
		fmt.Println("User controller create error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}