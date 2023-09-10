package main

import (
	"database/sql"
	"net/http"

	"github.com/joaomisturini/contabilista/src"
	"github.com/mattn/go-sqlite3"
)

const dbFile string = "contabilista.db"

func main() {
	_ = sqlite3.SQLiteConn{} // This is here bacause, somehow, it does not understand the driver as necessary

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	controller := src.NewController(db)

	http.HandleFunc("/", controller.GetRoot)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
