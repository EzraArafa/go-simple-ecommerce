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

func (r *CartRepository) GetCartItems() ([]models.CartItemResponse, error) {
	query := `SELECT cart_items.id, cart_items.product_id, products.name, products.price, cart_items.quantity
			FROM cart_items
			INNER JOIN products ON cart_items.product_id = products.id
			`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItemResponse

	for rows.Next() {
		var item models.CartItemResponse

		err := rows.Scan(&item.ID, &item.ProductID, &item.Name, &item.Price, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *CartRepository) DeleteCartItem(id int) error {
	query := `DELETE FROM cart_items WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("item keranjang tidak ditemukan")
	}
	return nil
}
