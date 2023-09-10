package src

import (
	"database/sql"
	"fmt"
	"time"
)

type BillService struct {
	db *sql.DB
}

func NewBillService(db *sql.DB) *BillService {
	return &BillService{
		db: db,
	}
}

func (s *BillService) CreateBill(request CreateBill) (BillModel, error) {
	model := NewBillModelFromRequest(request)

	repository := NewBillRepository(s.db)
	createdModel, err := repository.InsertBill(model)

	go simulateThirdParty(createdModel)

	return createdModel, err
}

func simulateThirdParty(model BillModel) {
	time.Sleep(8 * time.Second)

	fmt.Printf("Send bill # %d to third party", model.Id)
}
