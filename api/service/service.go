package service

import (
	"source/model"

	"gorm.io/gorm"
)

type ProductService interface {
	GetProduct(id int) (model.Tb_casestudy, error)
	GetAllProducts() ([]model.Tb_casestudy, error)
	CreateProduct(product model.Tb_casestudy) (model.Tb_casestudy, error)
	UpdateProduct(id int, product model.Tb_casestudy) (model.Tb_casestudy, error)
	DeleteProduct(id int) error
}

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{db: db}
}

func (s *productService) GetProduct(id int) (model.Tb_casestudy, error) {
	var product model.Tb_casestudy
	if err := s.db.First(&product, id).Error; err != nil {
		return model.Tb_casestudy{}, err
	}
	return product, nil
}

func (s *productService) GetAllProducts() ([]model.Tb_casestudy, error) {
	var products []model.Tb_casestudy
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) CreateProduct(product model.Tb_casestudy) (model.Tb_casestudy, error) {
	if err := s.db.Create(&product).Error; err != nil {

		return model.Tb_casestudy{}, err
	}
	return product, nil
}

func (s *productService) UpdateProduct(id int, updatedProduct model.Tb_casestudy) (model.Tb_casestudy, error) {
	var product model.Tb_casestudy
	if err := s.db.First(&product, id).Error; err != nil {
		return model.Tb_casestudy{}, err
	}
	product.Title = updatedProduct.Title
	product.Description = updatedProduct.Description
	product.Imageuri = updatedProduct.Imageuri
	product.Createddate = updatedProduct.Createddate
	s.db.Save(&product)
	return product, nil
}

func (s *productService) DeleteProduct(id int) error {
	return s.db.Delete(&model.Tb_casestudy{}, id).Error
}
