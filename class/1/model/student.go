package model

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	User  User   `json:"user"`
	Batch Batch  `json:"batch"`
}
