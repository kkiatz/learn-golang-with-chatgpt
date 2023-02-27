package main

import "fmt"

func main() {
	var x int = 5
	if x > 4 {
		fmt.Printf("%d bigger than 4\n", x)
	} else {
		fmt.Printf("%d smaller than 4\n", x)
	}

	var y = 3
	switch y {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5, 6:
		fmt.Println("multiple swift cases")
	default:
		fmt.Println("invalid number")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	// initialization and increment statement can be omitted
	fmt.Println("\ninitialization and increment statement can be omitted")
	i := 2
	for ; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	// infinite loop
	// for {}

	// Break statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("\nbreak")
			break
		}
	}

	// continue statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("\nskip number 5")
			continue
		}
		fmt.Printf("%d ", i)
	}
}
