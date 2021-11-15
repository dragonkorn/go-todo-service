package repository

import (
	"service/internal/config"
	"service/internal/infrastructure/database"
	"service/internal/model"
)

type ProductRepository struct {
	Config *config.Configuration
	Db     *database.DB
}

var productList []model.Product

func NewProductRepository(
	cfg *config.Configuration,
	db *database.DB,
) *ProductRepository {
	return &ProductRepository{
		Config: cfg,
		Db:     db,
	}
}

func (s *ProductRepository) Create(product *model.Product) (bool, error) {
	productList = append(productList, *product)

	result := s.Db.Connection.Create(product) // pass pointer of data to Create
	if result != nil {
		return true, nil
	}
	return false, nil
}

func (s *ProductRepository) List() (*[]model.Product, error) {
	return &productList, nil
}

func (s *ProductRepository) Get(id string) (*model.Product, error) {
	return nil, nil
}
