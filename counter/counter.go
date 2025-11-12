package main

import "fmt"

func counter() func() int {
	count := 10
	return func() int {
		for count > 0 { 
			count *= 2
			break
		}
		return count
	}
}

func main() {
	countFunc := counter()
	fmt.Println(countFunc())
	fmt.Println(countFunc())
	fmt.Println(countFunc())
}