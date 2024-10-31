package repository

import (
	"20241031/class/1/model"
	"database/sql"
)

type Student struct {
	Db *sql.DB
}

func CreateStudent(db *sql.DB) *Student {
	return &Student{Db: db}
}

func (r *Student) Create(student *model.Student) (*model.Student, error) {
	query := `INSERT INTO students (name,user_id, batch_id) VALUES ($1, $2, $3) RETURNING id`
	err := r.Db.QueryRow(query, student.Name, student.User.Id, student.Batch.Id).Scan(&student.Id)
	if err != nil {
		return nil, err
	}
	return student, nil
}
