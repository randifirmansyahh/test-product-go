package service

import (
	p "test-product/app/model/productModel"
	r "test-product/app/repository"
)

type service struct {
	repository r.IProductRepository
}

func NewService(repository r.IProductRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]p.Product, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(id int) (p.Product, error) {
	return s.repository.FindByID(id)
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
	return s.repository.Create(product)
}

func (s *service) Update(ID int, productRequest p.ProductRequest) (p.Product, error) {
	find, _ := s.repository.FindByID(ID)

	price, _ := productRequest.Price.Int64()
	rating, _ := productRequest.Rating.Int64()
	discount, _ := productRequest.Price.Int64()

	find.Title = productRequest.Title
	find.Price = int(price)
	find.Description = productRequest.Description
	find.Rating = int(rating)
	find.Discount = int(discount)

	return s.repository.Update(find)
}

func (s *service) Delete(ID int) (p.Product, error) {
	find, _ := s.repository.FindByID(ID)
	return s.repository.Delete(find)
}
