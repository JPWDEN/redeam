package service

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/redeam/go/data"
	"github.com/redeam/go/types"

	"github.com/bdlm/log"
)

//Server is the interface that defines the CRUD API object
type Server interface {
	HandleBooks(resp http.ResponseWriter, req *http.Request)
	AddBook(resp http.ResponseWriter, req *http.Request) error
	AdjustRating(adjust int, resp http.ResponseWriter, req *http.Request) error
	GetBook(titles []string, resp http.ResponseWriter, req *http.Request) error
	GetAllBooks(resp http.ResponseWriter, req *http.Request) error
	DeleteBook(title []string, resp http.ResponseWriter, req *http.Request) error
	ChangeStatus(checkIn bool, resp http.ResponseWriter, req *http.Request) error
	Collapse(resp http.ResponseWriter, req *http.Request) error
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
	//log.Infof("Path args: %+v, raw %s, length %d", pathArgs, req.URL.RawQuery, len(pathArgs))
	switch req.Method {
	case "GET":
		var err error
		if len(pathArgs) == 1 && !strings.Contains(req.URL.RawQuery, "title") {
			err = svr.GetAllBooks(resp, req)
		} else if strings.Contains(req.URL.RawQuery, "title") {
			err = svr.GetBook(m["title"], resp, req)
		}
		if err != nil {
			log.Errorf("GET Error: %v", err)
		}
		return
	case "POST":
		var err error
		switch pathArgs[1] {
		case "add":
			err = svr.AddBook(resp, req)
		case "up":
			err = svr.AdjustRating(1, resp, req)
		case "down":
			err = svr.AdjustRating(-1, resp, req)
		case "in":
			err = svr.ChangeStatus(true, resp, req)
		case "out":
			err = svr.ChangeStatus(false, resp, req)
		case "!":
			err = svr.Collapse(resp, req)
		}
		if err != nil {
			log.Errorf("POST Error: %v", err)
		}
		return
	case "DELETE":
		err := svr.DeleteBook(m["title"], resp, req)
		if err != nil {
			log.Errorf("DELETE Error: %v", err)
		}
		return
	default:
		respondHTTPErr(resp, req, http.StatusBadRequest)
		return
	}
}

//GetBook allows a user to request information on a given book entry
func (svr *ServerType) GetBook(title []string, resp http.ResponseWriter, req *http.Request) error {
	result, err := svr.DAO.SelectSingleBook(title)
	if err != nil {
		respondErr(resp, req, http.StatusInternalServerError, " Database error: ", err)
		return err
	}
	if result == nil {
		respond(resp, req, http.StatusNoContent, &result)
	} else {
		respond(resp, req, http.StatusOK, &result)
	}
	return nil
}

//GetAllBooks allows a user to request information on a given book entry
func (svr *ServerType) GetAllBooks(resp http.ResponseWriter, req *http.Request) error {
	result, err := svr.DAO.SelectAllBooks()
	if err != nil {
		respondErr(resp, req, http.StatusInternalServerError, " Database error: ", err)
		return err
	}
	if result == nil {
		respond(resp, req, http.StatusNoContent, &result)
	} else {
		respond(resp, req, http.StatusOK, &result)
	}
	return nil
}

//AddBook allows a user to add a new book to the database
//Multiple copies of a book are NOT currently allowed--Its a small bookshelf still
func (svr *ServerType) AddBook(resp http.ResponseWriter, req *http.Request) error {
	var book types.BookData
	err := decodeBody(req, &book)
	if err != nil {
		respondErr(resp, req, http.StatusBadRequest, " Failed to decode body: ", err)
		return err
	}
	if book.Title == "" {
		respondErr(resp, req, http.StatusBadRequest, " Missing body information: ", err)
		return nil
	}

	err = svr.DAO.InsertBook(book)
	if err != nil {
		respondErr(resp, req, http.StatusInternalServerError, " Database error: ", err)
		return err
	}
	respond(resp, req, http.StatusOK, &types.ListStatus{
		Status: "Success",
		Info:   fmt.Sprintf("Title '%s' added", book.Title),
	})
	return nil
}

//AdjustRating allows a user to add a new book to the database
func (svr *ServerType) AdjustRating(adjust int, resp http.ResponseWriter, req *http.Request) error {
	var book types.BookData
	err := decodeBody(req, &book)
	if err != nil {
		respondErr(resp, req, http.StatusBadRequest, "Failed to decode body: ", err)
	}

	err = svr.DAO.UpdateRating(book, adjust)
	if err != nil {
		respondErr(resp, req, http.StatusInternalServerError, " Database error: ", err)
		return err
	}
	respond(resp, req, http.StatusOK, &types.ListStatus{
		Status: "Success",
		Info:   fmt.Sprintf("Title '%s' rating updated", book.Title),
	})
	return nil
}

//ChangeStatus allows a user to change the state of a book:  Checked in or out
func (svr *ServerType) ChangeStatus(checkIn bool, resp http.ResponseWriter, req *http.Request) error {
	var book types.BookData
	err := decodeBody(req, &book)
	if err != nil {
		respondErr(resp, req, http.StatusBadRequest, "Failed to decode body: ", err)
	}

	newStatus, err := svr.DAO.UpdateStatus(book, checkIn)
	if err != nil {
		respondErr(resp, req, http.StatusInternalServerError, " Database error: ", err)
		return err
	}
	checkInStr := "Checked out"
	if newStatus {
		checkInStr = "Checked in"
	}
	respond(resp, req, http.StatusOK, &types.ListStatus{
		Status: "Success",
		Info:   fmt.Sprintf("Title '%s': %s", book.Title, checkInStr),
	})
	return nil
}

//DeleteBook allows an authenticated user to delete a given book entry
func (svr *ServerType) DeleteBook(title []string, resp http.ResponseWriter, req *http.Request) error {
	var book types.BookData
	err := decodeBody(req, &book)
	if err != nil {
		respondErr(resp, req, http.StatusBadRequest, "Failed to decode body: ", err)
		return err
	}

	err = svr.DAO.DeleteBookEntry(book)
	if err != nil {
		respondErr(resp, req, http.StatusInternalServerError, " Database error: ", err)
		return err
	}
	respond(resp, req, http.StatusOK, &types.ListStatus{
		Status: "Success",
		Info:   fmt.Sprintf("Title '%s' removed", book.Title),
	})
	return nil
}

//Its not production code, so lets take the opportunity to have fun!
func (svr *ServerType) Collapse(resp http.ResponseWriter, req *http.Request) error {
	diceBig1, err1 := rand.Int(rand.Reader, big.NewInt(100))
	diceBig2, err2 := rand.Int(rand.Reader, big.NewInt(100))
	if err1 != nil || err2 != nil {
		log.Error("No dice")
		respondHTTPErr(resp, req, http.StatusInternalServerError)
		return nil
	}
	if diceBig1.Int64() == diceBig2.Int64() {
		err := svr.DAO.Farenheit451()
		if err != nil {
			respondErr(resp, req, http.StatusInternalServerError, " Database error: ", err)
			return err
		}
		respond(resp, req, http.StatusResetContent, &types.ListStatus{Status: "Book List Truncated"})
	} else {
		respond(resp, req, http.StatusOK, &types.ListStatus{Status: "Book List Intact"})
	}
	return nil
}
