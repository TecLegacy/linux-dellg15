package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/teclegacy/golang-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAllProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	// if any error occurred during iteration of rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	query := `INSERT INTO products (id, name, description, image, price, quantity, createdAt) 
				VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := s.db.Exec(query, product.ID, product.Name, product.Description, product.Image, product.Price, product.Quantity, product.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeHolders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeHolders)

	//Convert ProductIDs to []interface{}
	args := make([]interface{}, len(productIDs))

	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	product := []types.Product{}

	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		product = append(product, *p)
	}

	return product, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.CreatedAt)
	if err != nil {
		return nil, err
	}

	return product, nil
}
