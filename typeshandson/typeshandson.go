package main

import "fmt"

type PaymentSystem interface {
	pay()
}

type CreditCard struct {
	creditCardNo int
	lastDate     float32
}

type UPI struct {
	amount     float32
	transnDate float32
}

type Cash struct {
	cashDenom int
	cashPin   float32
}

func (c CreditCard) pay() {
	fmt.Println("payment System : creditCard")
}

func (u UPI) pay() {
	fmt.Println("payment system : upi")
}

func (ch Cash) pay() {
	fmt.Println("payment system : cash")
}

func payment(payment PaymentSystem) {
	payment.pay()
}

func main() {
	creditCard1 := CreditCard{1234, 12.27}
	upi1 := UPI{230000.21, 30.09}
	cash1 := Cash{2000, 4111.32}

	payment(creditCard1)
	payment(upi1)
	payment(cash1)
}
