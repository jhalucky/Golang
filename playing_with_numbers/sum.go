package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	sum:= 1
	original:= num
	
	for num != 0 {
		digit := num % 10
		sum *= digit
		num /= 10
	}

	fmt.Printf("The sum of digits of %d is %d\n", original, sum)
}