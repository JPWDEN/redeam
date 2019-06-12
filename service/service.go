package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/redeam/data"
	"github.com/redeam/types"

	"github.com/bdlm/log"
)

//Server is the interface that defines the CRUD API object
type Server interface {
	HandleBooks(resp http.ResponseWriter, req *http.Request)
	BookChanges(resp http.ResponseWriter, req *http.Request) error
	GetBook(titles []string, resp http.ResponseWriter, req *http.Request) error
	GetAllBooks(resp http.ResponseWriter, req *http.Request) error
	UpdateBooks(book types.BookData, resp http.ResponseWriter, req *http.Request) error
	DeleteBooks(title []string, resp http.ResponseWriter, req *http.Request) error
}

//ServerType is the server object with db connection
type ServerType struct {
	DAO *data.StoreType
}

func encodeBody(resp http.ResponseWriter, req *http.Request, data interface{}) error {
	return json.NewEncoder(resp).Encode(data)
}

func decodeBody(req *http.Request, data interface{}) error {
	defer req.Body.Close()
	return json.NewDecoder(req.Body).Decode(data)
}

func respond(resp http.ResponseWriter, req *http.Request, status int, data interface{}) {
	resp.WriteHeader(status)
	if data != nil {
		encodeBody(resp, req, data)
	}
}

func respondErr(resp http.ResponseWriter, req *http.Request, status int, args ...interface{}) {
	respond(resp, req, status, map[string]interface{}{
		"error": map[string]interface{}{"message": fmt.Sprint(args...)},
	})
}

func respondHTTPErr(resp http.ResponseWriter, req *http.Request, status int) {
	respondErr(resp, req, status, http.StatusText(status))
}

//HandleBooks routes various API requests to the proper function
func (svr *ServerType) HandleBooks(resp http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	pathArgs := strings.Split(strings.Trim(path, "/"), "/")

	m, _ := url.ParseQuery(req.URL.RawQuery)
	log.Infof("query: %+v", m)
	log.Infof("Path args: %+v, URL: %+v, raw %s", pathArgs, path, req.URL.RawQuery)
	switch req.Method {
	case "GET":
		if len(pathArgs) == 1 && !strings.Contains(req.URL.RawQuery, "title") {
			err := svr.GetAllBooks(resp, req)
			if err != nil {

			}
		} else if strings.Contains(req.URL.RawQuery, "title") {
			err := svr.GetBook(m["title"], resp, req)
			if err != nil {

			}
		}
		return
	case "POST":
		err := svr.BookChanges(resp, req)
		if err != nil {

		}
		return
	case "DELETE":
		err := svr.DeleteBook(m["title"], resp, req)
		if err != nil {

		}
		return
	}
	respondHTTPErr(resp, req, http.StatusNotFound)
}

//GetBook allows a user to request information on a given book entry
func (svr *ServerType) GetBook(title []string, resp http.ResponseWriter, req *http.Request) error {
	result, err := svr.DAO.SelectSingleBook(title)
	if err != nil {

	}
	respond(resp, req, http.StatusOK, &result)
	return nil
}

//GetAllBooks allows a user to request information on a given book entry
func (svr *ServerType) GetAllBooks(resp http.ResponseWriter, req *http.Request) error {
	result, err := svr.DAO.SelectAllBooks()
	if err != nil {

	}
	respond(resp, req, http.StatusOK, &result)
	return nil
}

//BookChanges allows a user to update information about a given book entry
func (svr *ServerType) BookChanges(resp http.ResponseWriter, req *http.Request) error {
	var book types.BookData
	err := decodeBody(req, &book)
	if err != nil {
		respondErr(resp, req, http.StatusBadRequest, " Failed to read", err)
		return err
	}
	if book.Title == "" {
		respondErr(resp, req, http.StatusBadRequest, " Missing body information", err)
		return nil
	}
	log.Infof("Book to update: %+v", book)

	err = svr.DAO.CreateUpdateBook(book)
	if err != nil {

	}
	return nil
}

//DeleteBook allows an authenticated user to delete a given book entry
func (svr *ServerType) DeleteBook(title []string, resp http.ResponseWriter, req *http.Request) error {
	var book types.BookData
	err := decodeBody(req, &book)
	if err != nil {
		respondErr(resp, req, http.StatusBadRequest, "Failed to read", err)
	}
	log.Infof("Book to update: %+v", book)

	//err = svr.DAO.DeleteBookEntry()
	//if err != nil {
	//}

	return nil
}
