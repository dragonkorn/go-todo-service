package repository

import (
	"service/internal/config"
	"service/internal/model"
)

type ProductRepository struct {
	Config *config.Configuration
}

var productList []model.Product

func NewProductRepository(cfg *config.Configuration) *ProductRepository {
	return &ProductRepository{
		Config: cfg,
	}
}

func (s *ProductRepository) Create(product *model.Product) (bool, error) {
	productList = append(productList, *product)

	result := db.Create(product) // pass pointer of data to Create

	return true, nil
}

func (s *ProductRepository) List() (*[]model.Product, error) {
	return &productList, nil
}

func (s *ProductRepository) Get(id string) (*model.Product, error) {
	return nil, nil
}
