package types

type Money int64

type PaymentCategory string

type PaymentStatus string

const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

type Account struct {
	ID int64
	Phone Phone 
	Balance Money
}

type Phone string 

type Payment struct {
	ID int64
	Phone Phone
	Balance Money
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}
