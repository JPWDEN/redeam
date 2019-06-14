package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/redeam/go/client"
	"github.com/redeam/go/data"
	"github.com/redeam/go/service"
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

	db, err := sql.Open("mysql", "root@tcp(db:3306)/sys?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dao := &data.StoreType{DAO: db}
	svc := &service.ServerType{DAO: dao}

	go func() {
		if !test {
			return
		}
		for {
			time.Sleep(time.Minute)
			client.Run()
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/books/", svc.HandleBooks)
	log.Infof("Starting API on %s", addr)
	log.Fatal(http.ListenAndServe(port, mux))

	log.Info("Ending service")
}
