package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

//curl -v -X POST localhost:8080/books/add --data '{"title": "Redeam", "author": "JPW", "status": true}'
//Windows: curl -v -X POST localhost:8080/books/add --data "{\"title\": \"Redeam\", \"author\": \"JPW\", \"status\": true}"
func runAddBook() {
	route := "http://localhost:8080/books/add"
	payload := map[string]interface{}{
		"title":  "Redeam",
		"author": "JPW",
		"status": true,
	}
	byteMap, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	resp, err := http.Post(route, "application/json", bytes.NewBuffer(byteMap))
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Infof("result: %v", result)
}

//curl -vv -X POST localhost:8080/books/up --data '{"title": "Josh"}'
//curl -vv -X POST localhost:8080/books/down --data '{"title": "Josh"}'
//Windows: curl -vv -X POST localhost:8080/books/up --data "{\"title\": \"Josh\"}"
//Windows: curl -vv -X POST localhost:8080/books/down --data "{\"title\": \"Josh\"}"
func runAdjustRating() {
	//Hit up rating endpoint
	route := "http://localhost:8080/books/up"
	payload := map[string]interface{}{"title": "Redeam"}
	byteMap, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	resp, err := http.Post(route, "application/json", bytes.NewBuffer(byteMap))
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Infof("result: %v", result)

	//Hit down rating endpoint
	route = "http://localhost:8080/books/down"
	resp, err = http.Post(route, "application/json", bytes.NewBuffer(byteMap))
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Infof("result: %v", result)

	//Hit "user error"

}

//curl -v -X GET localhost:8080/books/?title="Joe"
func runGetBook() {
	url := "http://localhost:8080/books/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Infof("Error %v", err)
	}
	param := req.URL.Query()
	param.Add("title", "Joe")
	req.URL.RawQuery = param.Encode()
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

//curl -vv -X DELETE localhost:8080/books/ --data '{"title": "Josh"}'
//Windows: curl -vv -X DELETE localhost:8080/books/ --data "{\"title\": \"Josh\"}"
func runDeleteBook() {
	route := "http://localhost:8080/books/"
	payload := url.Values{}
	payload.Set("title", "Josh")
	req, err := http.NewRequest("DELETE", route, strings.NewReader(payload.Encode()))
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
	//Hit check-out endpoint
	route := "http://localhost:8080/books/out"
	payload := map[string]interface{}{"title": "Fire"}
	byteMap, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	resp, err := http.Post(route, "application/json", bytes.NewBuffer(byteMap))
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Infof("result: %v", result)

	//Hit check-in endpoint
	route = "http://localhost:8080/books/in"
	resp, err = http.Post(route, "application/json", bytes.NewBuffer(byteMap))
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Infof("result: %v", result)
}

//curl -vv -X POST localhost:8080/books/!
func runCollapse() {
	route := "http://localhost:8080/books/!"
	req, err := http.NewRequest("POST", route, nil)
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
	runAddBook()
	//runAdjustRating()
	//runGetBook()
	runGetAllBooks()
	//runDeleteBook()
	//runChangeStatus()
	//runCollapse()
}
