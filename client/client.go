package client

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//curl -v -X POST localhost:8080/books/add --data '{"title": "Redeam", "author": "JPW", "status": true}'
//Windows: curl -v -X POST localhost:8080/books/add --data "{\"title\": \"Redeam\", \"author\": \"JPW\", \"status\": true}"
func runAddBook() {
	url := ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Infof("Error %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Infof(string(body))
}

//curl -vv -X POST localhost:8080/books/up --data '{"title": "Josh"}'
//curl -vv -X POST localhost:8080/books/down --data '{"title": "Josh"}'
//Windows: curl -vv -X POST localhost:8080/books/up --data "{\"title\": \"Josh\"}"
//Windows: curl -vv -X POST localhost:8080/books/down --data "{\"title\": \"Josh\"}"
func runAdjustRating() {
	url := ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Infof("Error %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Infof(string(body))
}

//curl -v -X GET localhost:8080/books/?title="Joe"
func runGetBook() {
	url := ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Infof("Error %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Infof(string(body))
}

//curl -v -X GET localhost:8080/books/
func runGetAllBooks() {
	url := "http://localhost:8080/books/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Infof("Error %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Infof(string(body))
}

//curl -vv -X DELETE localhost:8080/books/up --data '{"title": "Josh"}'
//Windows: curl -vv -X DELETE localhost:8080/books/up --data "{\"title\": \"Josh\"}"
func runDeleteBook() {
	url := ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Infof("Error %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Infof(string(body))
}

//curl -vv -X POST localhost:8080/books/out --data '{"title": "Fire"}'
//curl -vv -X POST localhost:8080/books/in --data '{"title": "Fire"}'
//Windows: curl -vv -X POST localhost:8080/books/out --data "{\"title\": \"Fire\"}"
//Windows: curl -vv -X POST localhost:8080/books/in --data "{\"title\": \"Fire\"}"
func runChangeStatus() {
	url := ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Infof("Error %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Infof(string(body))
}

//curl -vv -X POST localhost:8080/books/!
func runCollapse() {
	url := ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Infof("Error %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Infof(string(body))
}

//Run performs curl calls for all subsequent endpoints
func Run() {
	//runAddBook()
	//runAdjustRating()
	//runGetBook()
	runGetAllBooks()
	//runDeleteBook()
	//runChangeStatus()
	//runCollapse()
}
