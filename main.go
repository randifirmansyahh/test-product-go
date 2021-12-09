package main

import (
	"log"
	"test-product/handler"
	"test-product/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/product_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	db.AutoMigrate(&product.Product{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()
	router.GET("/", productHandler.RootHandler)
	router.GET("/products", productHandler.GetProductsHandler)
	router.GET("/product/:id", productHandler.GetProductByIdHandler)
	router.POST("/product", productHandler.AddProductHandler)
	router.PUT("/product/:id", productHandler.UpdateProductHandler)
	router.DELETE("/product/:id", productHandler.DeleteProductHandler)
	router.Run()
}
