package repositories

import (
	"taskmanager/config"
	"taskmanager/models"
)

func CreateBook(Title string, Author string,Year int, Genre string, Available bool)error{
	query := `
	INSERT INTO books (title, author, year, genre, available)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := config.DB.Exec(query,
		Title,
		Author,
		Year,
		Genre,
		Available,
	)
	return err
}
func GetAllBooks()([]models.Book,error){
	query := "SELECT id, title, author, year, genre, available FROM books"

	rows,err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year, &book.Genre, &book.Available)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func GetBookByID(id int) (models.Book, error) {
	query := "SELECT id, title, author, year, genre, available FROM books WHERE id = $1"

	var book models.Book
	err := config.DB.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.Author, &book.Year, &book.Genre, &book.Available)

	return book, err
}
func DeleteBook(id int)(error){
	query := "DELETE FROM books WHERE id = $1"
	_,err := config.DB.Exec(query, id)
	return err
}
