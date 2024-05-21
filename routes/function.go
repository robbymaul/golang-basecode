package routes

import (
	"basecode/config"
	"basecode/controller"
	"basecode/middleware"
	"basecode/repository"
	"basecode/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	router := gin.Default()

	conf := config.Config{
		Dialect:  os.Getenv("DB_DIALECT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
	err := conf.GetConnection()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	conf.AutoMigration()

	middleware := middleware.Middleware{DB: config.DB}

	rgAuth := router.Group("/api/v1/auth")
	{
		repo := repository.NewUserRepository(config.DB)
		service := services.NewUserService(repo)
		controller := controller.NewAuthController(service)
		rgAuth.POST("/register", controller.Register)
		rgAuth.POST("/login", controller.Login)
	}

	rgProduct := router.Group("/api/v1/admin/products", middleware.Authentication(), middleware.CheckPermission("admin"))
	{
		repo := repository.NewProductRepository(config.DB)
		service := services.NewProductServiceImpl(repo)
		controller := controller.NewProductController(service)
		rgProduct.POST("/create", controller.CreateProduct)
	}

	return router
}
