// routes/api.go
package routes

import (
	"product-store/app/controllers"
	"product-store/app/repositories"
	"product-store/app/services"
	"product-store/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// group: v1
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		authController := controllers.AuthController{DB: db}
		v1.POST("/register", authController.Register)
		v1.POST("/login", authController.Login)

		// Product routes (protected)
		productRepo := repositories.NewProductRepository(db)
		productService := services.NewProductService(productRepo)
		productController := controllers.NewProductController(productService)

		productRoutes := v1.Group("/products")
		productRoutes.Use(middleware.AuthMiddleware())
		{
			productRoutes.GET("/", productController.GetProducts)
			productRoutes.POST("/", productController.CreateProduct)
		}
	}

	return router
}
