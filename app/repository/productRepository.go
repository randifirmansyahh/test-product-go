package repository

import (
	p "test-product/app/model/productModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]p.Product, error) {
	var products []p.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) FindByID(ID int) (p.Product, error) {
	var product p.Product
	err := r.db.Find(&product, ID).Error
	return product, err
}

func (r *repository) Create(product p.Product) (p.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) Update(product p.Product) (p.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) Delete(product p.Product) (p.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
