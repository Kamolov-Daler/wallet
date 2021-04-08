package types

// Money представляет собой денежную сумму в минималӣнҷх единиқах (центы, копейки, дирамы и т.д.).
type Money int64

// PaymentCategory представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type PaymentCategory string

// PaymentStatus представляет собой статус платежа.
type PaymentStatus string

// Предопредерённые статусы платежей.
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Payment представляет информацию о платеже.
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

// Phone телефон пользователя
type Phone string

// Account представляет информацию о счёте пользователя.
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
