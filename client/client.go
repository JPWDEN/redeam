package client

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

//curl -v -X POST localhost:8080/books/add --data '{"title": "Redeam", "author": "JPW", "status": true}'
//Windows: curl -v -X POST localhost:8080/books/add --data "{\"title\": \"Redeam\", \"author\": \"JPW\", \"status\": true}"
func runAddBook() {
	cmd, err := exec.Command("").CombinedOutput()
	if err != nil {
		log.Errorf("curl error: %v", err)
	}
	log.Infof("dir: %+v", string(cmd))

}

//curl -vv -X POST localhost:8080/books/up --data '{"title": "Josh"}'
//curl -vv -X POST localhost:8080/books/down --data '{"title": "Josh"}'
//Windows: curl -vv -X POST localhost:8080/books/up --data "{\"title\": \"Josh\"}"
//Windows: curl -vv -X POST localhost:8080/books/down --data "{\"title\": \"Josh\"}"
func runAdjustRating() {
	cmd, err := exec.Command("").CombinedOutput()
	if err != nil {
		log.Errorf("curl error: %v", err)
	}
	log.Infof("dir: %+v", string(cmd))
}

//curl -v -X GET localhost:8080/books/?title="Joe"
func runGetBook() {
	cmd, err := exec.Command("").CombinedOutput()
	if err != nil {
		log.Errorf("curl error: %v", err)
	}
	log.Infof("dir: %+v", string(cmd))
}

//curl -v -X GET localhost:8080/books/
func runGetAllBooks() {
	cmd, err := exec.Command("").CombinedOutput()
	if err != nil {
		log.Errorf("curl error: %v", err)
	}
	log.Infof("dir: %+v", string(cmd))
}

//curl -vv -X DELETE localhost:8080/books/up --data '{"title": "Josh"}'
//Windows: curl -vv -X DELETE localhost:8080/books/up --data "{\"title\": \"Josh\"}"
func runDeleteBook() {
	cmd, err := exec.Command("").CombinedOutput()
	if err != nil {
		log.Errorf("curl error: %v", err)
	}
	log.Infof("dir: %+v", string(cmd))
}

//curl -vv -X POST localhost:8080/books/out --data '{"title": "Fire"}'
//curl -vv -X POST localhost:8080/books/in --data '{"title": "Fire"}'
//Windows: curl -vv -X POST localhost:8080/books/out --data "{\"title\": \"Fire\"}"
//Windows: curl -vv -X POST localhost:8080/books/in --data "{\"title\": \"Fire\"}"
func runChangeStatus() {
	cmd, err := exec.Command("").CombinedOutput()
	if err != nil {
		log.Errorf("curl error: %v", err)
	}
	log.Infof("dir: %+v", string(cmd))
}

//curl -vv -X POST localhost:8080/books/!
func runCollapse() {
	cmd, err := exec.Command("").CombinedOutput()
	if err != nil {
		log.Errorf("curl error: %v", err)
	}
	log.Infof("dir: %+v", string(cmd))
}

//Run performs curl calls for all subsequent endpoints
func Run() {
	runAddBook()
	runAdjustRating()
	runGetBook()
	runGetAllBooks()
	runDeleteBook()
	runChangeStatus()
	runCollapse()
}
