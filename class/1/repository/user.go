package repository

import (
	"20241031/class/1/model"
	"database/sql"
)

type User struct {
	Db *sql.DB
}

func CreateUser(db *sql.DB) *User {
	return &User{Db: db}
}

func (r *User) Create(user *model.User) error {
	query := `INSERT INTO users (username,password) VALUES ($1, $2) RETURNING id`
	err := r.Db.QueryRow(query, user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *User) Authenticate(user *model.User) (*model.User, error) {
	query := `SELECT id, username, password, role FROM users WHERE username=$1 AND password=$2`
	var customerResponse model.User
	err := r.Db.QueryRow(query, user.Username, user.Password).Scan(&customerResponse.Id, &customerResponse.Username, &customerResponse.Password, &customerResponse.Role)

	if err != nil {
		return nil, err
	}

	return &customerResponse, nil
}
