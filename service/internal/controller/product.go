package controller

import (
	"fmt"
	"net/http"
	"service/internal/config"
	"service/internal/model"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Config         *config.Configuration
	ProductService *service.ProductService
}

func NewProductController(
	cfg *config.Configuration,
	ps *service.ProductService,
) *ProductController {
	return &ProductController{
		Config:         cfg,
		ProductService: ps,
	}
}

func (c *ProductController) Create(ctx *gin.Context) {
	var reqProduct model.Product
	if err := ctx.ShouldBindJSON(&reqProduct); err != nil {
		fmt.Println("Product controller create ShouldBindJSON error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("reqProduct : ", reqProduct)
	result, err := c.ProductService.Create(&reqProduct)

	if err != nil {
		fmt.Println("Product controller create error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  err,
			"result": result,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
	return
}

func (c *ProductController) List(ctx *gin.Context) {
	response, err := c.ProductService.List()
	if err != nil {
		fmt.Println("Product controller list error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  err,
			"result": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": response,
	})
}
