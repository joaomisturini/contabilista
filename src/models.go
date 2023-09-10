package src

type BillModel struct {
	Id          int64
	Description string
	Amount      int64
}

func NewBillModelFromRequest(request CreateBill) BillModel {
	return BillModel{
		Description: request.Description,
		Amount:      request.Amount,
	}
}
