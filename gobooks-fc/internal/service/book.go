package service

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	Id     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books(title, author, genre) VALUES (?, ?, ?)"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.Id = int(lastId)
	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "SELECT id, title, author, genre FROM books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Genre); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *BookService) GetBookById(id int) (*Book, error) {
	query := "SELECT id, title, author, genre FROM books WHERE id=?"
	row := s.db.QueryRow(query, id)
	var book Book
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) UpdateBook(book *Book, id int) error {
	query := "UPDATE books SET title=?, author=?, genre=? WHERE id=?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, id)
	return err
}

func (s *BookService) DeleteBook(id int) error {
	query := "DELETE from books where id=?"
	_, err := s.db.Exec(query, id)
	return err
}

func (s *BookService) SearchBooksByName(name string) ([]Book, error) {
	query := "SELECT id, title, author, genre from books where title like ?"
	rows, err := s.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *BookService) SimulateReading(bookId int, duration time.Duration, results chan<- string) {
	book, err := s.GetBookById(bookId)
	if err != nil || book == nil {
		results <- fmt.Sprintf("Book %d not found", bookId)
		return
	}
	time.Sleep(duration)
	results <- fmt.Sprintf("Book %s read", book.Title)
}

func (s *BookService) SimulateMultipleReadings(bookIds []int, duration time.Duration) []string {
	results := make(chan string, len(bookIds))

	for _, id := range bookIds {
		go func(bookId int) {
			s.SimulateReading(bookId, duration, results)
		}(id)
	}

	var responses []string
	for range bookIds {
		responses = append(responses, <-results)
	}
	close(results)
	return responses
}
