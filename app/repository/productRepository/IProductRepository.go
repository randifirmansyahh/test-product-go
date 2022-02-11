package productRepository

import "test-product/app/model/productModel"

type IProductRepository interface {
	FindAll() ([]productModel.Product, error)
	FindByID(ID int) (productModel.Product, error)
	Create(product productModel.Product) (productModel.Product, error)
	Update(product productModel.Product) (productModel.Product, error)
	Delete(product productModel.Product) (productModel.Product, error)
}
