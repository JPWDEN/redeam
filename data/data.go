package data

import (
	"database/sql"

	"github.com/bdlm/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redeam/types"
)

//Store is the interface defining the object for db functions
type Store interface {
	SQLTest(book types.BookData) error
	SelectAllBooks() ([]types.BookData, error)
	SelectSingleBook(option []string) ([]types.BookData, error)
	CreateUpdateBook(book types.BookData) error
	DeleteBookEntry() error
}

//StoreType is the struct holding the db connection
type StoreType struct {
	DAO *sql.DB
}

//SQLTest tests the db connection
func (store *StoreType) SQLTest(book types.BookData) error {
	_, err := store.DAO.Exec(`
INSERT INTO books VALUES (?, ?, ?, ?, ?, ?)`, book.Title, book.Author, book.Publisher, book.PublishDate, book.Rating, book.Status)
	if err != nil {
		log.Error(err)
	}
	return nil
}

//SelectAllBooks selects a user-chosen column for a range of books
func (store *StoreType) SelectAllBooks() ([]types.BookData, error) {
	results, err := store.DAO.Query(`SELECT title, author, rating, status FROM books`)
	if err != nil {
		log.Errorf("Error querying mysql: %v", err)
		return nil, err
	}
	var tags []types.BookData
	for results.Next() {
		var tag types.BookData
		err = results.Scan(&tag.Title, &tag.Author, &tag.Rating, &tag.Status)
		if err != nil {
			log.Warnf("Error selecting single row: %v", err)
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

//SelectSingleBook selects a single book from the db table
func (store *StoreType) SelectSingleBook(option []string) ([]types.BookData, error) {
	results, err := store.DAO.Query(`SELECT author, rating, status FROM books WHERE title = ?`, option[0])
	if err != nil {
		log.Errorf("Error querying mysql: %v", err)
		return nil, err
	}
	var tags []types.BookData
	for results.Next() {
		var tag types.BookData
		err = results.Scan(&tag.Author, &tag.Rating, &tag.Status)
		if err != nil {
			log.Warnf("Error selecting single row: %v", err)
			return nil, err
		}
		tag.Title = option[0]
		tags = append(tags, tag)
	}
	return tags, nil
}

//CreateUpdateBook updates a book.  If it doesn't exist, it is added to the db
func (store *StoreType) CreateUpdateBook(book types.BookData) error {
	var count int
	err := store.DAO.QueryRow(`SELECT count(*) FROM books WHERE title = ?`, book.Title).Scan(&count)
	if err != nil {

	}
	if count == 0 {
		_, err := store.DAO.Exec(`INSERT INTO books VALUES (?, ?, ?, ?, ?, ?)`, book.Title, book.Author, book.Publisher, book.PublishDate, book.Rating, book.Status)
		if err != nil {

		}

	} else {
		_, err := store.DAO.Exec(`UPDATE books SET`)
		if err != nil {

		}

	}

	return nil
}

//DeleteBookEntry removes a book from the database
func (store *StoreType) DeleteBookEntry(title string) error {
	_, err := store.DAO.Exec(`DELETE FROM books WHERE title = ?`, title)
	if err != nil {

	}
	return nil
}
