package service

import (
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/redeam/go/data"
)

func TestGetAllBooks(test *testing.T) {
	var (
	//expected = []string{}
	)

	db, mock, err := sqlmock.New()
	if err != nil {
		test.Errorf("Error creating sql mock: %v", err)
	}
	dao := &data.StoreType{DAO: db}
	server := ServerType{DAO: dao}

	mock.ExpectBegin()
	//mock.ExpectedQuery(``).WithArgs(1).WillReturnRows(sqlmock.NewRows(columns))

	var resp http.ResponseWriter
	var req *http.Request
	err = server.GetAllBooks(resp, req)
	if err != nil {

	}

	err = db.Close()
	if err != nil {

	}
}
