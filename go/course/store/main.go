package main

import (
	"course/store/customer"
	"course/store/order"
	"course/store/product"
	"encoding/json"
	"fmt"
)

func main() {
	product := product.Product{
		Name:        "Processador Intel Core I7 3.7GHZ",
		Description: "Este processador é exclusivo para jogos digitais",
		Price:       237.90,
		Stock:       200,
	}

	customer := customer.Customer{
		Name:     "Murilo Locatelli",
		Email:    "murilolocatelli@gmail.com",
		Password: "123456",
	}

	order := order.CreateOrder(product, customer, 5)

	fmt.Println(convertToJSON(order))
	fmt.Println(order.Customer.Name)
}

func convertToJSON(s interface{}) string {
	j, _ := json.Marshal(s)

	return string(j)
}
