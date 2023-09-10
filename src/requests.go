package src

import "fmt"

type CreateBill struct {
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

func (r *CreateBill) IsValid() error {
	if r.Description == "" {
		return fmt.Errorf("description must not be empty")
	}

	if r.Amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	return nil
}
