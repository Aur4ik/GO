package models

type Book struct {
	Id       int    `json:"Id"`
	Title    string `json:"Title"`
	Author   string `json:"Author"`
	Year     int    `json:"Year"`
	Genre    string `json:"Genre"`
	Available bool   `json:"Avalible"`
}
