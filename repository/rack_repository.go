package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"strconv"

	"go.uber.org/zap"
)

type RackRepository interface {
	GetAll(page, limit int) ([]model.Rack, error)
	GetByID(id int) (*model.Rack, error)
	Create(rack model.Rack) (int, error)
	Update(rack model.Rack) error
	Delete(id int) error
	CountAll() (int, error)
}

type rackRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewRackRepository(db *sql.DB, log *zap.Logger) RackRepository {
	return &rackRepository{
		db:  db,
		log: log,
	}
}

func (r *rackRepository) GetAll(page, limit int) ([]model.Rack, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query("SELECT id, name, description FROM racks ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var racks []model.Rack
	for rows.Next() {
		var rs model.Rack
		if err := rows.Scan(&rs.ID, &rs.Name, &rs.Description); err != nil {
			r.log.Error("error : ", zap.Error(err))
			return nil, err
		}
		racks = append(racks, rs)
	}
	return racks, nil
}

func (r *rackRepository) GetByID(id int) (*model.Rack, error) {
	row := r.db.QueryRow("SELECT id, name, description FROM racks WHERE id = $1", id)
	var rs model.Rack
	if err := row.Scan(&rs.ID, &rs.Name, &rs.Description); err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	return &rs, nil
}

func (r *rackRepository) Create(rack model.Rack) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO racks(name, description) VALUES($1, $2) RETURNING id",
		rack.Name, rack.Description).Scan(&id)
	r.log.Info("rack created : ", zap.String("insert", "id : "+strconv.Itoa(id)))
	return id, err
}

func (r *rackRepository) Update(rack model.Rack) error {
	_, err := r.db.Exec(
		"UPDATE racks SET name=$1, description=$2 WHERE id=$3",
		rack.Name, rack.Description, rack.ID)
	r.log.Info("rack updated : ", zap.String("update", "id : "+strconv.Itoa(rack.ID)))
	return err
}

func (r *rackRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM racks WHERE id = $1", id)
	r.log.Info("rack deleted : ", zap.String("delete", "id : "+strconv.Itoa(id)))
	return err
}

func (r *rackRepository) CountAll() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM racks").Scan(&count)
	return count, err
}
