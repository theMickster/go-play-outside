package models

import "fmt"

type Order struct {
	ProductCode int
	Quantity    int
	Status      OrderStatus
}

type OrderStatus int

const (
	none OrderStatus = iota
	new
	received
	reserved
	filled
)

type InvalidOrder struct {
	Order Order
	Err   error
}

func (o Order) String() string {
	return fmt.Sprintf("Product Code: %v  --  Quantity: %v -- Status: %v", o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(status OrderStatus) string {
	switch status {
	case none:
		return "None"
	case new:
		return "New"
	case received:
		return "Received"
	case reserved:
		return "Reserved"
	case filled:
		return "Filled"
	default:
		return "Unknown Status"
	}
}
