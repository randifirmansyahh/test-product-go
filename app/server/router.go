package server

import (
	"os"
	"test-product/app/handler"
	"test-product/app/helper"
	"test-product/app/model/productModel"
	"test-product/app/repository"
	"test-product/app/repository/productRepository"
	"test-product/app/service/productService"

	"test-product/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_DATABASE_NAME")), &gorm.Config{})

	helper.CekConnectionDB(err)

	db.AutoMigrate(&productModel.Product{})

	productRepository := repository.Repository{
		ProductRepository: productRepository.NewRepository(db),
	}

	productService := service.Service{
		ProductService: productService.NewService(productRepository),
	}

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
