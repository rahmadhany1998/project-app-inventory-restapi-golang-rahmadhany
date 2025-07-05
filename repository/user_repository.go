package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
)

type UserRepository interface {
	GetAll(page, limit int) ([]model.User, error)
	GetByID(id int) (*model.User, error)
	Create(user model.User) (int, error)
	Update(user model.User) error
	Delete(id int) error
	CountAll() (int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll(page, limit int) ([]model.User, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query("SELECT id, name, email, role, password, status FROM users ORDER BY id ASC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.Password, &u.Status); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, role, password, status FROM users WHERE id = $1", id)
	var u model.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.Password, &u.Status); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Create(user model.User) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO users(name, email, role, password, status) VALUES($1, $2, $3, $4, $5) RETURNING id",
		user.Name, user.Email, user.Role, user.Password, user.Status).Scan(&id)
	return id, err
}

func (r *userRepository) Update(user model.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET name=$1, email=$2, role=$3, password=$4, status=$5 WHERE id=$6",
		user.Name, user.Email, user.Role, user.Password, user.Status, user.ID)
	return err
}

func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *userRepository) CountAll() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	return count, err
}
