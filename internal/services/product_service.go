package services

import (
	"errors"

	"github.com/EzraArafa/go-simple-ecommerce/internal/models"
	"github.com/EzraArafa/go-simple-ecommerce/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("nama produk tidak boleh kosong")
	}
	if product.Price <= 0 {
		return errors.New("harga produk harus lebih besar dari 0")
	}
	if product.Stock < 0 {
		return errors.New("stok produk tidak boleh minus")
	}

	err := s.repo.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil

}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	if products == nil {
		products = []models.Product{}
	}

	return products, nil
}
