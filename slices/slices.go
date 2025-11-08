package main 

import "fmt"

func main() {
	var nums = make([]int, 10)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums = append(nums, 10)
	fmt.Println(nums)
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:",cap(nums))
}