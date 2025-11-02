package main 

import "fmt"

func main() {
	var input string
	fmt.Print("Enter the string: ")
    fmt.Scan(&input)

	reversed := ""
	for i:=len(input)-1; i >= 0;i-- {
		reversed += string(input[i])
	}

	if input == reversed {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
}
}