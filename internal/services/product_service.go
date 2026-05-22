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

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(id int, product models.Product) error {
	if product.Price <= 0 {
		return errors.New("Harga produk harus lebih dari 0")
	}

	if product.Stock < 0 {
		return errors.New("Stok produk tidak boleh kurang dari 0")
	}

	err := s.repo.UpdateProduct(id, product)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
