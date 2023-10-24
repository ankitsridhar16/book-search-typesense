package postgres

import (
	"database/sql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

const (
	fetchAllBooksQuery = "SELECT * FROM books"
)

type DB struct {
	*sql.DB
}

type Book struct {
	ID              int
	Title           string
	Authors         []string
	PublicationYear int
	AverageRating   float64
	ImageURL        string
	RatingsCount    int
}

// Init initialize postgres connection
func Init(connStr string) (*DB, error) {
	db, dbErr := sql.Open("postgres", connStr)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	return &DB{db}, nil
}

// FetchDataFromDB fetches data from DB based on query
func (db *DB) FetchDataFromDB() ([]Book, error) {
	// fetch all rows from DB
	rows, dbErr := db.Query(fetchAllBooksQuery)
	if dbErr != nil {
		log.Fatal(dbErr)
		return nil, dbErr
	}
	defer rows.Close()

	// Load results to Book Struct
	var books []Book
	for rows.Next() {
		var book Book
		var authors pq.StringArray // omit authors in Book Struct
		scanErr := rows.Scan(&book.ID, &book.Title, &authors, &book.PublicationYear, &book.AverageRating,
			&book.ImageURL, &book.RatingsCount)
		if scanErr != nil {
			return nil, scanErr
		}

		books = append(books, book)
	}

	return books, nil
}
