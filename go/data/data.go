package data

import (
	"database/sql"
	"errors"

	"github.com/bdlm/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redeam/go/types"
)

//Store is the interface defining the object for db functions
type Store interface {
	SQLTest(book types.BookData) error
	SelectAllBooks() ([]types.BookData, error)
	SelectSingleBook(option []string) ([]types.BookData, error)
	UpdateRating(book types.BookData, adjust int) error
	UpdateStatus(book types.BookData, checkIn bool) (bool, error)
	InsertBook(book types.BookData) error
	DeleteBookEntry(book types.BookData) error
	Farenheit451() error
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

//UpdateRating updates the rating of a title in the database
func (store *StoreType) UpdateRating(book types.BookData, adjust int) error {
	var count int
	err := store.DAO.QueryRow(`SELECT count(*) FROM books WHERE title = ?`, book.Title).Scan(&count)
	if count == 0 {
		return errors.New("Title does not exist")
	}

	_, err = store.DAO.Exec(`UPDATE books SET rating = rating + ? WHERE title = ?`, adjust, book.Title)
	if err != nil {
		log.Errorf("Error updating rating: %v", err)
		return err
	}
	return nil
}

//UpdateStatus updates the status of a book:  Is it checkin in or out?
func (store *StoreType) UpdateStatus(book types.BookData, checkIn bool) (bool, error) {
	var status bool
	err := store.DAO.QueryRow(`SELECT status FROM books WHERE title = ?`, book.Title).Scan(&status)
	if err != nil {
		log.Errorf("Error querying mysql: %v", err)
		return status, err
	}

	if status != checkIn {
		_, err := store.DAO.Exec(`UPDATE books SET status = ? WHERE title = ?`, checkIn, book.Title)
		if err != nil {
			log.Errorf("Error updating rating: %v", err)
			return false, err
		}
	}
	return checkIn, nil
}

//InsertBook updates a book.  If it doesn't exist, it is added to the db
func (store *StoreType) InsertBook(book types.BookData) error {
	var count int
	err := store.DAO.QueryRow(`SELECT count(*) FROM books WHERE title = ?`, book.Title).Scan(&count)
	if err != nil {
		log.Errorf("Error selecting book: %v", err)
		return err
	}
	if count == 0 {
		_, err := store.DAO.Exec(`INSERT INTO books VALUES (?, ?, ?, ?, ?, ?)`, book.Title, book.Author, book.Publisher, book.PublishDate, book.Rating, book.Status)
		if err != nil {
			log.Errorf("Error inserting book: %v", err)
			return err
		}

	} else {
		log.Infof("Book %s already exists; returning", book.Title)
	}
	return nil
}

//DeleteBookEntry removes a book from the database
func (store *StoreType) DeleteBookEntry(book types.BookData) error {
	_, err := store.DAO.Exec(`DELETE FROM books WHERE title = ?`, book.Title)
	return err
}

//There was a fire and everything was lost...Time to start fresh!
func (store *StoreType) Farenheit451() error {
	_, err := store.DAO.Exec(`TRUNCATE TABLE books`)
	return err
}
