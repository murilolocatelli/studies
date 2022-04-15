package order

import (
	"course/store/customer"
	"course/store/product"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// ENUMS
const (
	CREATED = iota + 1
	AWATING
	SHIPPING
)

// Order struct
type Order struct {
	Number   string
	Status   int
	Product  product.Product
	Customer customer.Customer
}

// CreateOrder func
func CreateOrder(product product.Product, customer customer.Customer, amount int) Order {

	product.Stock -= amount
	product.Price *= float32(amount)

	order := Order{
		Number:   generateNumber(6),
		Status:   CREATED,
		Product:  product,
		Customer: customer,
	}

	return order
}

func generateNumber(amount int) string {
	s1 := rand.NewSource(time.Now().UnixNano())

	var nums string

	for count := 0; count < amount; count++ {
		number := rand.New(s1).Intn(100)
		nums += strconv.Itoa(number)
	}

	return fmt.Sprintf("ML-%s", nums)
}
