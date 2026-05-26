package services

import (
	"errors"

	"github.com/EzraArafa/go-simple-ecommerce/internal/models"
	"github.com/EzraArafa/go-simple-ecommerce/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddToCart(item models.CartItem) error {
	if item.Quantity <= 0 {
		return errors.New("jumlah barang minimal 1")
	}
	if item.ProductID <= 0 {
		return errors.New("ID produk tidak valid")
	}
	return s.repo.AddToCart(item)
}
