package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/redeam/client"
	"github.com/redeam/data"
	"github.com/redeam/service"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, rd *http.Request) {
		wr.Header().Set("Access-Control-Allow-Origin", "*")
		//wr.Header().Set("Access-Control-Expose-Headers")
		fn(wr, rd)
	}
}

func main() {
	log.Info("Starting main service")

	port := os.Getenv("PORT")
	addr := os.Getenv("ADDRESS")
	if port == "" || addr == "" {
		port = ":8080"
		addr = "localhost"
	}
	test := false
	if os.Getenv("TEST") == "" || os.Getenv("TEST") == "true" {
		test = true
	}
	log.Infof("Test %t", test)

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/sys?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dao := &data.StoreType{DAO: db}

	go func() {
		//if !test {
		//	return
		//}
		for {
			time.Sleep(time.Second * 5)
			client.Run()
		}
	}()

	svc := &service.ServerType{DAO: dao}
	mux := http.NewServeMux()
	mux.HandleFunc("/books/", svc.HandleBooks)
	log.Infof("Starting API on %s", addr)
	log.Fatal(http.ListenAndServe(port, mux))

	log.Info("Ending service")
}
