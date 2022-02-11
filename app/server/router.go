package server

import (
	"test-product/app/handler"
	"test-product/app/helper"
	pModel "test-product/app/model/productModel"
	pRepository "test-product/app/repository"
	pService "test-product/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {

	db, err := gorm.Open(mysql.Open(GetConnectionString()), &gorm.Config{})

	helper.CekConnectionDB(err)

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
