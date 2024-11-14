// app/services/product_service.go
package services

import (
	"product-store/app/models"
	"product-store/app/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetAllProducts(page, limit int, sort string, order string) ([]models.Product, int64, error) {
	offset := (page - 1) * limit
	return s.Repo.GetAll(offset, limit, sort, order)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.Create(product)
}
