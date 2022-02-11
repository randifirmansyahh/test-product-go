package repository

import p "test-product/app/model"

type IProductRepository interface {
	FindAll() ([]p.Product, error)
	FindByID(ID int) (p.Product, error)
	Create(product p.Product) (p.Product, error)
	Update(product p.Product) (p.Product, error)
	Delete(product p.Product) (p.Product, error)
}
