package main 

import "fmt"

func main() {
    
	var i int
	
	fmt.Print("Enter the number: ")
	fmt.Scan(&i)

	if i % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}