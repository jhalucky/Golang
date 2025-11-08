package main 

import "fmt"

func main() {

	m:= make(map[string]string)

	m["name"] = "Lucky"
	m["Languages"] = "English, Hindi, Go, Python"

	fmt.Println(m["Languages"])

	j:= make(map[string]int)

	j["age"] = 25
	j["experience"] = 5
    
	fmt.Println(j["age"])
}