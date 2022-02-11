package productService

import (
	p "test-product/app/model/productModel"
	r "test-product/app/repository"
)

type service struct {
	repository r.Repository
}

func NewService(repository r.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]p.Product, error) {
	return s.repository.ProductRepository.FindAll()
}

func (s *service) FindByID(id int) (p.Product, error) {
	return s.repository.ProductRepository.FindByID(id)
}

func (s *service) Create(productRequest p.ProductRequest) (p.Product, error) {
	price, _ := productRequest.Price.Int64()
	rating, _ := productRequest.Rating.Int64()
	discount, _ := productRequest.Price.Int64()
	product := p.Product{
		Title:       productRequest.Title,
		Price:       int(price),
		Description: productRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	return s.repository.ProductRepository.Create(product)
}

func (s *service) Update(ID int, productRequest p.ProductRequest) (p.Product, error) {
	find, _ := s.repository.ProductRepository.FindByID(ID)

	price, _ := productRequest.Price.Int64()
	rating, _ := productRequest.Rating.Int64()
	discount, _ := productRequest.Price.Int64()

	find.Title = productRequest.Title
	find.Price = int(price)
	find.Description = productRequest.Description
	find.Rating = int(rating)
	find.Discount = int(discount)

	return s.repository.ProductRepository.Update(find)
}

func (s *service) Delete(ID int) (p.Product, error) {
	find, _ := s.repository.ProductRepository.FindByID(ID)
	return s.repository.ProductRepository.Delete(find)
}
