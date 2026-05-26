package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/EzraArafa/go-simple-ecommerce/internal/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddToCart(item models.CartItem) error {
	query := `INSERT INTO cart_items (product_id, quantity) VALUES (?, ?)`

	_, err := r.db.Exec(query, item.ProductID, item.Quantity)
	if err != nil {
		if strings.Contains(err.Error(), "foreign key constraint fails") {
			return errors.New("produk tidak ditemukan")
		}
		return err
	}
	return nil
}
