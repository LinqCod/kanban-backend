package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/kanban-backend/internal/model"
)

const (
	GetAllProductsQuery = `SELECT id, title, price, description, category, rate FROM products;`
)

type ProductRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewProductRepository(ctx context.Context, db *sql.DB) *ProductRepository {
	return &ProductRepository{
		ctx: ctx,
		db:  db,
	}
}

func (u ProductRepository) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product

	rows, err := u.db.QueryContext(u.ctx, GetAllProductsQuery)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(
			&product.Id,
			&product.Title,
			&product.Price,
			&product.Description,
			&product.Category,
			&product.Rate,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
