package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	nums = append(nums, 4)

	fmt.Println(nums)
}
