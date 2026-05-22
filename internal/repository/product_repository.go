package repository

import (
	"database/sql"
	"errors"

	"github.com/EzraArafa/go-simple-ecommerce/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	query := `
		INSERT INTO products (name, description, price, stock)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = int(id)

	return nil
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product

		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = ?`

	var p models.Product

	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Produk tidak ditemukan")
		}
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) UpdateProduct(id int, product models.Product) error {
	query := `UPDATE products
			  SET name = ?, description = ?, price = ?, stock = ?, updated_at = CURRENT_TIMESTAMP
			  WHERE id = ?`

	result, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}
