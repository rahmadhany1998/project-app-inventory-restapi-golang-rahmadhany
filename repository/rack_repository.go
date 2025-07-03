package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
)

type RackRepository interface {
	GetAll() ([]model.Rack, error)
	GetByID(id int) (*model.Rack, error)
	Create(rack model.Rack) (int, error)
	Update(rack model.Rack) error
	Delete(id int) error
}

type rackRepository struct {
	db *sql.DB
}

func NewRackRepository(db *sql.DB) RackRepository {
	return &rackRepository{db}
}

func (r *rackRepository) GetAll() ([]model.Rack, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM racks ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var racks []model.Rack
	for rows.Next() {
		var rs model.Rack
		if err := rows.Scan(&rs.ID, &rs.Name, &rs.Description); err != nil {
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
		return nil, err
	}
	return &rs, nil
}

func (r *rackRepository) Create(rack model.Rack) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO racks(name, description) VALUES($1, $2) RETURNING id",
		rack.Name, rack.Description).Scan(&id)
	return id, err
}

func (r *rackRepository) Update(rack model.Rack) error {
	_, err := r.db.Exec(
		"UPDATE racks SET name=$1, description=$2 WHERE id=$3",
		rack.Name, rack.Description, rack.ID)
	return err
}

func (r *rackRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM racks WHERE id = $1", id)
	return err
}
