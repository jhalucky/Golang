package main 

import "fmt"

func main() {
	var input string
	fmt.Print("ENter a string: ")
	fmt.Scan(&input)
    
	reversed:= ""
	for i:=len(input)-1;i>=0;i--{
        reversed += string(input[i]) 
	}

	fmt.Println("Reversed String: ", reversed)
}