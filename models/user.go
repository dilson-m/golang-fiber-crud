package models

type User struct {
	Id        uint   `json "id"`
	FirstName string `json "nome"`
	LastName  string `json "sobrenome"`
	Email     string `json "email"`
	Password  string `json "senha"`
}
