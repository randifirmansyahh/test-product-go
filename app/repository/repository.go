package repository

import "test-product/app/repository/productRepository"

type Repository struct {
	ProductRepository productRepository.IProductRepository
}
