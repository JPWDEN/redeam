package types

import "github.com/go-sql-driver/mysql"

//Books is the JSON-relatable object used for API calls
type BookData struct {
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Publisher   string         `json:"publisher"`
	PublishDate mysql.NullTime `json:"publish_date"`
	Rating      int            `json:"rating"`
	Status      bool           `json:"status"`
}

type ListStatus struct {
	Status string `json:"status"`
	Info string `json:"info"`
}
