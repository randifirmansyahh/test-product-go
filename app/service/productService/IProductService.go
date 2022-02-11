package productService

import p "test-product/app/model/productModel"

type ProductService interface {
	FindAll() ([]p.Product, error)
	FindByID(ID int) (p.Product, error)
	Create(productRequest p.ProductRequest) (p.Product, error)
	Update(ID int, productRequest p.ProductRequest) (p.Product, error)
	Delete(ID int) (p.Product, error)
}
