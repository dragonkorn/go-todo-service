package container

import (
	"fmt"
	"service/internal/config"
	"service/internal/controller"
	"service/internal/infrastructure/http"
	"service/internal/repository"
	"service/internal/service"

	"go.uber.org/dig"
)

type Container struct {
	Container *dig.Container
}

func (c *Container) Configuration() error {
	c.Container = dig.New()

	if err := c.Container.Provide(config.NewConfiguration); err != nil {
		fmt.Println(err)
		return err
	}
	if err := c.Container.Provide(http.NewHttpServer); err != nil {
		fmt.Println(err)
		return err
	}

	// Controller
	if err := c.Container.Provide(controller.NewProductController); err != nil {
		fmt.Println(err)
		return err
	}

	// Service
	if err := c.Container.Provide(service.NewProductService); err != nil {
		fmt.Println(err)
		return err
	}

	// Repository
	if err := c.Container.Provide(repository.NewProductRepository); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *Container) Run() error {
	fmt.Println("Run Container")

	if err := c.Container.Invoke(func(s *http.HttpServer) {
		s.Run()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	c := &Container{}
	if err := c.Configuration(); err != nil {
		return nil, err
	}

	return c, nil
}
