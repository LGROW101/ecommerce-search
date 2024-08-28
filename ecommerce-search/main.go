package main

import (
	"log"
	"time"

	"github.com/LGROW101/ecommerce-search/config"
	"github.com/LGROW101/ecommerce-search/database"
	"github.com/LGROW101/ecommerce-search/handler"
	"github.com/LGROW101/ecommerce-search/repository"
	"github.com/LGROW101/ecommerce-search/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	var db *gorm.DB
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = database.InitDB(cfg)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retrying in 5 seconds... (Attempt %d/%d)", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to initialize database after %d attempts: %v", maxRetries, err)
	}

	productRepo := repository.NewProductRepository(db)
	searchService := service.NewSearchService(productRepo)
	searchHandler := handler.NewSearchHandler(searchService)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/search", searchHandler.Search)
	e.GET("/products/:id", searchHandler.GetProductDetails)

	e.Logger.Fatal(e.Start(":8080"))
}
