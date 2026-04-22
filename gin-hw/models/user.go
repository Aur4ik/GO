package models

type User struct{
	Id int 			`json:"Id"`
	Email string	`json:"Email"`
	Password string	`json:"Password"`
}