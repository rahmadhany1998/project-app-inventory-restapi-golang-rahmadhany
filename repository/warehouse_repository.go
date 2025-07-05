package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"strconv"

	"go.uber.org/zap"
)

type WarehouseRepository interface {
	GetAll(page, limit int) ([]model.Warehouse, error)
	GetByID(id int) (*model.Warehouse, error)
	Create(warehouse model.Warehouse) (int, error)
	Update(warehouse model.Warehouse) error
	Delete(id int) error
	CountAll() (int, error)
}

type warehouseRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewWarehouseRepository(db *sql.DB, log *zap.Logger) WarehouseRepository {
	return &warehouseRepository{
		db:  db,
		log: log,
	}
}

func (r *warehouseRepository) GetAll(page, limit int) ([]model.Warehouse, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query("SELECT id, name, description FROM warehouses ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var warehouses []model.Warehouse
	for rows.Next() {
		var w model.Warehouse
		if err := rows.Scan(&w.ID, &w.Name, &w.Description); err != nil {
			r.log.Error("error : ", zap.Error(err))
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
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	return &w, nil
}

func (r *warehouseRepository) Create(warehouse model.Warehouse) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO warehouses(name, description) VALUES($1, $2) RETURNING id",
		warehouse.Name, warehouse.Description).Scan(&id)
	r.log.Info("warehouse created : ", zap.String("insert", "id : "+strconv.Itoa(id)))
	return id, err
}

func (r *warehouseRepository) Update(warehouse model.Warehouse) error {
	_, err := r.db.Exec(
		"UPDATE warehouses SET name=$1, description=$2 WHERE id=$3",
		warehouse.Name, warehouse.Description, warehouse.ID)
	r.log.Info("warehouse updated : ", zap.String("update", "id : "+strconv.Itoa(warehouse.ID)))
	return err
}

func (r *warehouseRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM warehouses WHERE id = $1", id)
	r.log.Info("warehouse deleted : ", zap.String("delete", "id : "+strconv.Itoa(id)))
	return err
}

func (r *warehouseRepository) CountAll() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM warehouses").Scan(&count)
	return count, err
}
