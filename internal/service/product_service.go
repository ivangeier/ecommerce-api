package service

import (
	"github.com/ivangeier/ig-commerce/ecommerce-api/internal/database"
	"github.com/ivangeier/ig-commerce/ecommerce-api/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductByCategoryId(categoryId string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name string, description string, price float64, categoryId string, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryId, imageURL)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
