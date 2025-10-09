package model

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID        string `json:"id" binding:"len=0"` // validasi len=0 digunakan supaya user tidak dapat input id sendiri
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	IsDeleted *bool  `json:"is_deleted,omitempty"`
	// menggunakan pointer karena nantinya data tidak dikembalikan ke user / tidak ditampilkan
}

var (
	errDBNil = errors.New("koneksi db tidak tersedia")
)

// Select all product
func SelectProducts(db *sql.DB) ([]Product, error) {
	if db == nil {
		return nil, errDBNil
	}

	query := `SELECT id, name, price FROM products WHERE is_deleted = false;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	products := []Product{}
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// Select product by ID
func SelectProductById(db *sql.DB, id string) (Product, error) {
	if db == nil {
		return Product{}, errDBNil
	}

	var product Product
	query := `SELECT id, name, price FROM products WHERE is_deleted = false AND id = $1;`
	row := db.QueryRow(query, id)
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

// Create product
func InsertProduct(db *sql.DB, product Product) error {
	if db == nil {
		return errDBNil
	}

	query := `INSERT INTO products (id, name, price) VALUES ($1, $2, $3);`
	_, err := db.Exec(query, product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(db *sql.DB, product Product) error {
	if db == nil {
		return errDBNil
	}

	query := `UPDATE products SET name=$1, price=$2 WHERE id=$3;`

	_, err := db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete Product
func DeleteProduct(db *sql.DB, id string) error {
	if db == nil {
		return errDBNil
	}

	query := `DELETE FROM products WHERE id = $1;`
	_, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
