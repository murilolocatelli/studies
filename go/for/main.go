package main

import "fmt"

func soma(nums []int) int {
	var count = 0

	for _, num := range nums {
		count += num
	}

	return count
}

func main() {
	nums := []int{1, 2, 3}

	resultado := soma(nums)

	fmt.Println(resultado)
}
