package service

import (
	"fmt"
	"service/internal/config"
	"service/internal/model"
	"service/internal/repository"
)

type ProductService struct {
	Config            *config.Configuration
	ProductRepository *repository.ProductRepository
}

func NewProductService(
	cfg *config.Configuration,
	productRepo *repository.ProductRepository,
) *ProductService {
	return &ProductService{
		Config:            cfg,
		ProductRepository: productRepo,
	}
}

func (s *ProductService) Create(product *model.Product) (bool, error) {
	result, err := s.ProductRepository.Create(product)
	if err != nil {
		fmt.Println("Product Service Create Received err", err)
		return false, err
	}
	return result, nil
}

func (s *ProductService) List() (*[]model.Product, error) {
	productList, err := s.ProductRepository.List()
	if err != nil {
		fmt.Println("Product Service List Received err", err)
		return nil, err
	}
	return productList, nil
}

func (s *ProductService) Get(id uint) (*model.Product, error) {
	return nil, nil
}
