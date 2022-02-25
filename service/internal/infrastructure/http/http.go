package http

import (
	"fmt"
	"service/internal/config"
	"service/internal/controller"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Config      *config.Configuration
	route       *gin.Engine
	ProductCtrl *controller.ProductController
	UserCtrl *controller.UserController
}

func NewHttpServer(
	cfg *config.Configuration,
	productCtrl *controller.ProductController,
	userCtrl *controller.UserController,
) *HttpServer {
	fmt.Println("http  new server called")
	r := gin.New()
	return &HttpServer{
		Config:      cfg,
		route:       r,
		ProductCtrl: productCtrl,
		UserCtrl: userCtrl,
	}
}

func (s *HttpServer) Configuration() *gin.Engine {
	r := s.route

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	product := r.Group("/product")
	{
		product.POST("/", s.ProductCtrl.Create)
		product.GET("/", s.ProductCtrl.List)
	}
	user := r.Group("/user")
	{
		user.POST("/", s.UserCtrl.CreateUser)
	}

	return r
}

func (s *HttpServer) Run() error {
	r := s.Configuration()

	if err := r.Run(fmt.Sprintf(":%s", s.Config.Port)); err != nil {
		return err
	}

	return nil
}
