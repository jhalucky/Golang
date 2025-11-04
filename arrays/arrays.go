package main 

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	var name [4]string
	name[0]="Alice"
	name[1]="Hardick"
	name[2]="Aman"
	fmt.Println(name)

	var scores = [5]int{90, 80, 70, 60, 50}
	fmt.Println(scores)
}