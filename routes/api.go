// routes/api.go
package routes

import (
	"product-store/app/controllers"
	"product-store/app/repositories"
	"product-store/app/services"
	"product-store/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	// ALLOWED_ORIGIN := os.Getenv("ALLOWED_ORIGIN")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		authController := controllers.AuthController{DB: db}
		v1.POST("/register", authController.Register)
		v1.POST("/login", authController.Login)
		v1.POST("/logout", authController.Logout)

		userController := controllers.UserController{DB: db}
		v1.GET("/user", middleware.AuthMiddleware(), userController.GetUserByToken)

		// Product routes (protected)
		productRepo := repositories.NewProductRepository(db)
		productService := services.NewProductService(productRepo)
		productController := controllers.NewProductController(productService)

		productRoutes := v1.Group("/products")
		// productRoutes.Use(middleware.AuthMiddleware())
		{
			productRoutes.POST("/get", productController.GetProducts)
			productRoutes.POST("/create", productController.CreateProduct)
		}
	}

	return router
}
