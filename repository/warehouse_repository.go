package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
)

type WarehouseRepository interface {
	GetAll() ([]model.Warehouse, error)
	GetByID(id int) (*model.Warehouse, error)
	Create(warehouse model.Warehouse) (int, error)
	Update(warehouse model.Warehouse) error
	Delete(id int) error
}

type warehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepository{db}
}

func (r *warehouseRepository) GetAll() ([]model.Warehouse, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM warehouses ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []model.Warehouse
	for rows.Next() {
		var w model.Warehouse
		if err := rows.Scan(&w.ID, &w.Name, &w.Description); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}
	return warehouses, nil
}

func (r *warehouseRepository) GetByID(id int) (*model.Warehouse, error) {
	row := r.db.QueryRow("SELECT id, name, description FROM warehouses WHERE id = $1", id)
	var w model.Warehouse
	if err := row.Scan(&w.ID, &w.Name, &w.Description); err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *warehouseRepository) Create(warehouse model.Warehouse) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO warehouses(name, description) VALUES($1, $2) RETURNING id",
		warehouse.Name, warehouse.Description).Scan(&id)
	return id, err
}

func (r *warehouseRepository) Update(warehouse model.Warehouse) error {
	_, err := r.db.Exec(
		"UPDATE warehouses SET name=$1, description=$2 WHERE id=$3",
		warehouse.Name, warehouse.Description, warehouse.ID)
	return err
}

func (r *warehouseRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM warehouses WHERE id = $1", id)
	return err
}
