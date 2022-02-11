package router

import (
	"log"
	handler "test-product/app/handler"
	pModel "test-product/app/model"
	pRepository "test-product/app/repository"
	gConnection "test-product/app/server"
	pService "test-product/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {

	db, err := gorm.Open(mysql.Open(gConnection.GetConnectionString()), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	db.AutoMigrate(&pModel.Product{})

	productRepository := pRepository.NewRepository(db)
	productService := pService.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	route := gin.Default()
	route.GET("/", productHandler.RootHandler)
	route.GET("/products", productHandler.GetProductsHandler)
	route.GET("/product/:id", productHandler.GetProductByIdHandler)
	route.POST("/product", productHandler.AddProductHandler)
	route.PUT("/product/:id", productHandler.UpdateProductHandler)
	route.DELETE("/product/:id", productHandler.DeleteProductHandler)
	route.Run()
}
