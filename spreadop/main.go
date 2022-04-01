package main

import "fmt"

func main() {
	nums := make([]int, 3)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	fmt.Println(nums)

	other := []int{4, 5}

	nums = append(nums, other...)

	fmt.Println(nums)
}
