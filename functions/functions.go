package main 

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main () {
	fmt.Println(getLanguages())
}

func getLanguages() (string, string) {
	return "Go", "Python"
}

// in golang, return functions can return multiple values