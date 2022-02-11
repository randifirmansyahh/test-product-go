package handler

import (
	"fmt"
	"net/http"
	"strconv"

	productModel "test-product/app/model"
	product "test-product/app/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is running",
	})
}

func (h *productHandler) GetProductsHandler(c *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errors": err,
		})
		return
	}

	var productsResponse []productModel.ProductResponse
	for _, p := range products {
		result := convertToProductResponse(p)
		productsResponse = append(productsResponse, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productsResponse,
	})
}

func (h *productHandler) GetProductByIdHandler(c *gin.Context) {
	params := c.Param("id")
	id, _ := strconv.Atoi(params)
	getProduct, err := h.productService.FindByID(id)

	if err != nil || getProduct.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"errors": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": convertToProductResponse(getProduct)})
}

func (h *productHandler) AddProductHandler(c *gin.Context) {
	var productRequest productModel.ProductRequest
	err := c.ShouldBindJSON(&productRequest)

	if err != nil {
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		return
	}

	product, err := h.productService.Create(productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Posted",
		"data":    convertToProductResponse(product),
	})
}

func (h *productHandler) UpdateProductHandler(c *gin.Context) {
	var productRequest productModel.ProductRequest
	err := c.ShouldBindJSON(&productRequest)

	if err != nil {
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		return
	}

	params := c.Param("id")
	id, _ := strconv.Atoi(params)
	product, err := h.productService.Update(id, productRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated",
		"data":    convertToProductResponse(product),
	})
}

func (h *productHandler) DeleteProductHandler(c *gin.Context) {
	params := c.Param("id")
	id, _ := strconv.Atoi(params)
	getProduct, err := h.productService.Delete(id)

	if err != nil || getProduct.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"errors": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "has been deleted"})
}

func convertToProductResponse(Obj productModel.Product) productModel.ProductResponse {
	return productModel.ProductResponse{
		ID:          Obj.ID,
		Title:       Obj.Title,
		Price:       Obj.Price,
		Description: Obj.Description,
		Rating:      Obj.Rating,
		Discount:    Obj.Discount,
	}
}
