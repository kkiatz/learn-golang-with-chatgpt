package main

import "fmt"

func main() {
	fmt.Println("Example for new function")
	// Allocate a new zeroed value of type *int
	ptr := new(int)

	// Print the value of the pointer
	fmt.Println("Pointer value:", ptr)

	// Print the value of the pointed-to integer
	fmt.Println("Integer value:", *ptr)

	// Set the value of the pointed-to integer
	*ptr = 42

	// Print the new value of the pointed-to integer
	fmt.Println("New integer value:", *ptr)

	fmt.Println("Example for make function")
	/*
		In Go, you use the make built-in function to create and initialize three types of composite values: slices, maps, and channels.

		You should use make to create slices, maps, and channels because they are all implemented
	*/

	// Create a slice of length 3 and capacity 5
	slice := make([]int, 3, 5)
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice ", slice)

	// create a map with capacity of 3
	m := make(map[string]int, 3)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println("map ", m)

	/*
		When creating a slice, map, or channel in Go using the built-in make function, you can specify an optional capacity parameter. This parameter specifies the initial capacity of the data structure and can have a number of advantages:

		Performance improvement: Specifying the capacity up front can improve performance by reducing the number of reallocations needed as the data structure grows. This is because Go's built-in data structures are implemented as arrays under the hood, and when an array's capacity is exceeded, a new, larger array must be created and the existing values copied over. By pre-allocating the required capacity, you can minimize the number of these expensive copy operations.

		Memory management: By specifying the capacity of a data structure up front, you can reduce memory fragmentation and make more efficient use of memory. This is because when Go's garbage collector runs, it tries to reclaim memory in contiguous blocks. If your data structure has been growing and shrinking without a specified capacity, it can be more difficult for the garbage collector to find a contiguous block of memory to reclaim.

		Simpler code: Specifying the capacity up front can simplify your code by reducing the need for reallocation and resizing operations. If you know that your data structure will be large enough to hold a certain number of items, you can simply allocate the required capacity up front and avoid any resizing operations later on.

		In Go, length and capacity are properties of a slice, map, or channel that describe its current size and capacity, respectively.

		For a slice, the length is the number of elements currently in the slice, and the capacity is the maximum number of elements that the slice can hold without allocating more memory. The length of a slice can be obtained using the built-in len function, and the capacity can be obtained using the built-in cap function.
	*/
	mySlice := make([]int, 0, 5)
	fmt.Println(len(mySlice)) // Output: 0
	fmt.Println(cap(mySlice)) // Output: 5

	mySlice = append(mySlice, 1)
	fmt.Println(len(mySlice)) // Output: 1
	fmt.Println(cap(mySlice)) // Output: 5

	mySlice = append(mySlice, 2, 3, 4, 5)
	fmt.Println(len(mySlice)) // Output: 5
	fmt.Println(cap(mySlice)) // Output: 5

}
