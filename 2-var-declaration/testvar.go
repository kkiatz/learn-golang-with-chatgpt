package main

import "fmt"

func main() {
	var a int = 1
	b := 1
	c := a + b
	fmt.Println("result of a + b is", c)
	result := fmt.Sprintf("result with sprintf is %d", c)
	fmt.Println(result)
	fmt.Println("Test string concatenantion")

	firstName := "Ali"
	lastName := "Ahmad"
	userName := firstName + " " + lastName
	fmt.Println(userName)

}
