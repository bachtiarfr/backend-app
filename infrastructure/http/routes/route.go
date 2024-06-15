package routes

import (
	"backend-app/application/services"
	"backend-app/config"
	"backend-app/domain/repositories"
	"backend-app/infrastructure/http/handlers"
	"backend-app/infrastructure/http/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	swipeRepo := repositories.NewSwipeRepository(db)
	swipeService := services.NewSwipeService(swipeRepo)
	swipeHandler := handlers.NewSwipeHandler(swipeService)

	api := router.Group("/api")
	{
		api.POST("/signup", userHandler.SignUp)
		api.POST("/login", userHandler.Login)

		auth := api.Group("/auth")
		auth.Use(middlewares.AuthMiddleware(cfg.SecretKey))
		{
			auth.POST("/swipe/:profileID/:action", swipeHandler.Swipe)
			auth.POST("/purchase", userHandler.PurchasePremium)
		}
	}
}
