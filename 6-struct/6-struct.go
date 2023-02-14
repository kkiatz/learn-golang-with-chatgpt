package main

import (
	"fmt"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Foo", Age: 1}
	fmt.Println(person)
	fmt.Println("Name ", person.Name)
	fmt.Println("Age ", person.Age)

	type Customer struct {
		Name string
		Age  int
	}

	var cust Customer
	cust.Name = "Ali"
	cust.Age = 3

	fmt.Println(cust)
	fmt.Println("Name ", cust.Name)
	fmt.Println("Age ", cust.Age)
	fmt.Printf("Print Name and age field %+v\n", cust)

}
