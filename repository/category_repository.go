package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
	"strconv"

	"go.uber.org/zap"
)

type CategoryRepository interface {
	GetAll(page, limit int) ([]model.Category, error)
	GetByID(id int) (*model.Category, error)
	Create(category model.Category) (int, error)
	Update(category model.Category) error
	Delete(id int) error
	CountAll() (int, error)
}

type categoryRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewCategoryRepository(db *sql.DB, log *zap.Logger) CategoryRepository {
	return &categoryRepository{
		db:  db,
		log: log,
	}
}

func (r *categoryRepository) GetAll(page, limit int) ([]model.Category, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query("SELECT id, name, description FROM categories ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			r.log.Error("error : ", zap.Error(err))
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *categoryRepository) GetByID(id int) (*model.Category, error) {
	row := r.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id)
	var c model.Category
	if err := row.Scan(&c.ID, &c.Name, &c.Description); err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	return &c, nil
}

func (r *categoryRepository) Create(category model.Category) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO categories(name, description) VALUES($1, $2) RETURNING id",
		category.Name, category.Description).Scan(&id)
	r.log.Info("category created : ", zap.String("insert", "id : "+strconv.Itoa(id)))
	return id, err
}

func (r *categoryRepository) Update(category model.Category) error {
	_, err := r.db.Exec(
		"UPDATE categories SET name=$1, description=$2 WHERE id=$3",
		category.Name, category.Description, category.ID)
	r.log.Info("category updated : ", zap.String("update", "id : "+strconv.Itoa(category.ID)))
	return err
}

func (r *categoryRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id = $1", id)
	r.log.Info("category deleted : ", zap.String("delete", "id : "+strconv.Itoa(id)))
	return err
}

func (r *categoryRepository) CountAll() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	return count, err
}
