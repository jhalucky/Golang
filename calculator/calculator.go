package main 

import "fmt"

func main() {
	var a, b float64
	var op string
	var result float64

	fmt.Print("Enter a: ")
	fmt.Scan(&a)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Enter b: ")
	fmt.Scan(&b)

    switch op {
    case "+":
	   result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
    case "/":
		result = a/b
	default:
		fmt.Println("Invalid Operator!")
		return
}

    fmt.Println("The result is: ", result)


	
	// fmt.Println("Sum of a and b is:", a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a/b)
	
}