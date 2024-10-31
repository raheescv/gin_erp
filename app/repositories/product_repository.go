// app/repositories/product_repository.go
package repositories

import (
	"product-store/app/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) GetAll(offset int, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var count int64

	repo.DB.Model(&models.Product{}).Count(&count)
	err := repo.DB.Offset(offset).Limit(limit).Find(&products).Error
	return products, count, err
}

func (repo *ProductRepository) Create(product *models.Product) error {
	return repo.DB.Create(product).Error
}
