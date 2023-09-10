package src

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Controller struct {
	db *sql.DB
}

func NewController(db *sql.DB) *Controller {
	return &Controller{
		db: db,
	}
}

func (c *Controller) GetRoot(w http.ResponseWriter, r *http.Request) {
	var data CreateBill

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := data.IsValid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := NewBillService(c.db)
	bill, err := service.CreateBill(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("got / request")
	io.WriteString(w, fmt.Sprintf("This is my website! New ID: %d", bill.Id))
}
