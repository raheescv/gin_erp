// routes/web.go
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
	r := gin.Default()

	// Auth routes
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// Product routes (protected)
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	productRoutes := r.Group("/products")
	productRoutes.Use(middleware.AuthMiddleware())
	{
		productRoutes.GET("/", productController.GetProducts)
		productRoutes.POST("/", productController.CreateProduct)
	}

	return r
}
