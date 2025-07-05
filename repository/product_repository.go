package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"strconv"

	"go.uber.org/zap"
)

type ProductRepository interface {
	GetAll(page, limit int) ([]model.Product, error)
	GetByID(id int) (*model.Product, error)
	Create(product model.Product) (int, error)
	Update(product model.Product) error
	Delete(id int) error
	CountAll() (int, error)
}

type productRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewProductRepository(db *sql.DB, log *zap.Logger) ProductRepository {
	return &productRepository{
		db:  db,
		log: log,
	}
}

func (r *productRepository) GetAll(page, limit int) ([]model.Product, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query("SELECT id, name, category_id, rack_id, warehouse_id, inventory_count, retail_price, selling_price, image FROM products ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.RackID, &p.WarehouseID, &p.InventoryCount, &p.RetailPrice, &p.SellingPrice, &p.Image); err != nil {
			r.log.Error("error : ", zap.Error(err))
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *productRepository) GetByID(id int) (*model.Product, error) {
	row := r.db.QueryRow("SELECT id, name, category_id, rack_id, warehouse_id, inventory_count, retail_price, selling_price, image FROM products WHERE id = $1", id)
	var p model.Product
	if err := row.Scan(&p.ID, &p.Name, &p.CategoryID, &p.RackID, &p.WarehouseID, &p.InventoryCount, &p.RetailPrice, &p.SellingPrice, &p.Image); err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) Create(product model.Product) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO products(name, category_id, rack_id, warehouse_id, inventory_count, retail_price, selling_price, image) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		product.Name, product.CategoryID, product.RackID, product.WarehouseID, product.InventoryCount, product.RetailPrice, product.SellingPrice, product.Image).Scan(&id)
	r.log.Info("product created : ", zap.String("insert", "id : "+strconv.Itoa(id)))
	return id, err
}

func (r *productRepository) Update(product model.Product) error {
	_, err := r.db.Exec(
		"UPDATE products SET name=$1, category_id=$2, rack_id=$3, warehouse_id=$4, inventory_count=$5, retail_price=$6, selling_price=$7, image=$8 WHERE id=$9",
		product.Name, product.CategoryID, product.RackID, product.WarehouseID, product.InventoryCount, product.RetailPrice, product.SellingPrice, product.Image, product.ID)
	r.log.Info("product updated : ", zap.String("update", "id : "+strconv.Itoa(product.ID)))
	return err
}

func (r *productRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	r.log.Info("product deleted : ", zap.String("delete", "id : "+strconv.Itoa(id)))
	return err
}

func (r *productRepository) CountAll() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	return count, err
}
