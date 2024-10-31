// models/product.go
package models

type Product struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}
