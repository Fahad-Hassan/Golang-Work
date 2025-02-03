package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32)
}
type payment struct {
	gateway paymenter
}

// open close principle (open for extension but close for modification)
func (p payment) makePayment(amount float32) {
	// razorpayPayment := razorpay{}
	// razorpayPayment.pay(amount)
	// stripepayPayment := stripe{}
	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

func (r razorpay) refund(amount float32) {
	fmt.Println("making refund payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)

}
func (r stripe) refund(amount float32) {
	fmt.Println("making refund payment using stripe", amount)
}

type paypal struct{}

func (r paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}
func (r paypal) refund(amount float32) {
	fmt.Println("making refund payment using paypal", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fakepayment", amount)
}
func (r fakepayment) refund(amount float32) {
	fmt.Println("making refund payment using fakepayment", amount)
}

func main() {
	// stripepayPaymentGW := stripe{}
	// razorpayPaymentGW := razorpay{}
	fakeGW := fakepayment{}
	newpayment := payment{
		gateway: fakeGW,
	}
	newpayment.makePayment(100)
}
