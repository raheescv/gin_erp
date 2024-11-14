// app/repositories/product_repository.go
package repositories

import (
	"fmt"
	"product-store/app/models"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) GetAll(offset int, limit int, sort string, order string) ([]models.Product, int64, error) {
	var products []models.Product
	var count int64

	// Validate sort column to prevent SQL injection
	validSortColumns := map[string]bool{"id": true, "name": true, "price": true}
	if !validSortColumns[sort] {
		sort = "id"
	}

	// Validate order to prevent SQL injection
	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	// Apply sorting and pagination
	repo.DB.Model(&models.Product{}).Count(&count)
	err := repo.DB.Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", sort, order)).Find(&products).Error
	print(count)
	print(err)
	return products, count, err
}

func (repo *ProductRepository) Create(product *models.Product) error {
	return repo.DB.Create(product).Error
}
