package main

import (
	"fmt"
	"link-nest/configs"
	"link-nest/internal/api"
	"link-nest/internal/auth"
	"link-nest/internal/database"
	"link-nest/internal/repository"
	"link-nest/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load Configuration
	cfg := configs.LoadConfig()

	// 2. Initialize Database
	db := database.InitDB(cfg)

	// 3. Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	// TODO: Initialize other repositories when needed
	// templateRepo := repository.NewTemplateRepository(db)
	// userPageRepo := repository.NewUserPageRepository(db)
	// moduleRepo := repository.NewModuleRepository(db)
	// paymentRepo := repository.NewPaymentRepository(db)

	// 4. Initialize Services
	userService := service.NewUserService(userRepo)
	// TODO: Initialize other services when needed
	// templateService := service.NewTemplateService(templateRepo)
	// userPageService := service.NewUserPageService(userPageRepo)
	// moduleService := service.NewModuleService(moduleRepo)
	// paymentService := service.NewPaymentService(paymentRepo)

	// 5. Initialize API Handlers
	userHandler := api.NewUserHandler(userService)

	// 6. Setup Gin Router
	router := gin.Default()

	// Public routes (no authentication required)
	publicRoutes := router.Group("/api/v1")
	{
		publicRoutes.POST("/register", userHandler.Register)
		publicRoutes.POST("/login", userHandler.Login)
	}

	// Authenticated routes
	authRoutes := router.Group("/api/v1")
	authRoutes.Use(auth.AuthMiddleware())
	{
		// User routes will be added here
		// Template routes will be added here
		// UserPage routes will be added here
		// Module routes will be added here
		// Payment routes will be added here
	}

	fmt.Printf("Server is running on %s\n", cfg.Server.Port)
	log.Fatal(router.Run(cfg.Server.Port))
}
