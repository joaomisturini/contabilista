package src

import "database/sql"

type BillRepository struct {
	db *sql.DB
}

func NewBillRepository(db *sql.DB) *BillRepository {
	return &BillRepository{
		db: db,
	}
}

func (r *BillRepository) InsertBill(model BillModel) (BillModel, error) {
	query := "INSERT INTO bills (description, amount) VALUES (?, ?)"

	result, err := r.db.Exec(query, model.Id, model.Description, model.Amount)
	if err != nil {
		return model, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return model, err
	}

	model.Id = lastId

	return model, nil
}
