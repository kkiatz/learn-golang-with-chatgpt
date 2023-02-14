package main

import (
	"fmt"
)

func main() {
	/*
		Go provides a number of built-in collection types, including arrays, slices, maps, and channels.

		Arrays: An array is a fixed-size sequence of elements of the same type. In Go, arrays are declared using the syntax var name [size]type, where name is the name of the array, size is the number of elements in the array, and type is the type of the elements in the array. For example: var arr [3]int.

		Slices: A slice is a dynamic array that can grow or shrink as needed. In Go, slices are declared using the syntax var name []type, where name is the name of the slice, and type is the type of the elements in the slice. Slices are created from arrays or other slices using the make function, like so: s := make([]int, 5).

		Maps: A map is a collection of key-value pairs. In Go, maps are declared using the syntax var name map[keyType]valueType, where name is the name of the map, keyType is the type of the keys, and valueType is the type of the values. Maps are created using the make function, like so: m := make(map[string]int).

		Channels: A channel is a typed conduit for communication between concurrent Go routines. In Go, channels are declared using the syntax var name chan<-/->type, where name is the name of the channel, type is the type of the elements in the channel, and <- specifies the direction of data flow (left arrow for receiving, right arrow for sending). Channels are created using the make function, like so: ch := make(chan int).

		In addition to these built-in collection types, Go also supports third-party libraries that provide additional collection types, such as sets, lists, and queues.
	*/
	fmt.Println("Example for Arrays")
	fmt.Println("Example for Arrays declaration 1")
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[2])

	fmt.Println("Example for Arrays declaration 2")
	var arr2 [2]int
	arr2[0] = 1
	arr2[1] = 2
	fmt.Println(arr2[1])

	fmt.Println("Example for Slices")
	// Declare a slice of integers
	var s []int

	// Initialize the slice with values using a slice literal
	s = []int{1, 2, 3, 4, 5}

	// Append a new value to the slice
	s = append(s, 6)

	// Get a sub-slice of the first three elements
	subSlice := s[:3]

	// Iterate over the slice using a for loop
	fmt.Println("Loop for slice of integers")
	for i, val := range s {
		fmt.Printf("Index %d has value %d\n", i, val)
	}

	// s[start:stop]
	fmt.Println("Loop for subslice of integers till 3rd element")
	for i, val := range subSlice {
		fmt.Printf("Index %d has value %d\n", i, val)
	}

	// s[start:stop]
	subSlice = s[1:3]
	fmt.Println("Loop for subslice of integers ")
	for i, val := range subSlice {
		fmt.Printf("Index %d has value %d\n", i, val)
	}

	// get the last element
	fmt.Println("Get the last element of s", s[len(s)-1])

	// prepend slices
	s = append([]int{100}, s...)
	fmt.Println("prepend slices s with 100", s)

	// insert slices
	slice := []int{1, 2, 3, 4}
	newElement := 5
	mid := len(slice) / 2
	newSlice := append(slice[:mid], append([]int{newElement}, slice[mid:]...)...)
	fmt.Println("Original slice", slice)
	fmt.Println("Insert newElement value 5 to middle of array", newSlice)

	fmt.Println("Example for Maps")
	m := map[string]int{"one": 1, "two": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// update key one
	m["one"] = 2
	fmt.Println("Update key one to 2", m)

	// delete key
	delete(m, "one")
	fmt.Println("Delete key one", m)

	// add new key
	m["three"] = 3
	fmt.Println("Add new key with value 3", m)

	// handle key error
	if _, ok := m["four"]; !ok {
		fmt.Println("Handle key error successfully")
	}

	fmt.Println("Example for Channels")
	/*
		Channels in Go are used to communicate and synchronize between goroutines, which are lightweight threads of execution that run concurrently. Channels provide a way for goroutines to send and receive data with each other, and they ensure that the communication is safe and synchronized.

		Here are some use cases for channels in Go:

		Coordination between goroutines: Channels can be used to coordinate the execution of multiple goroutines. For example, you can use a channel to signal to a goroutine that it should start executing, or to signal that it should stop.

		Sharing data between goroutines: Channels can be used to share data between multiple goroutines. For example, you can use a channel to pass a message from one goroutine to another, or to send a piece of data from one goroutine to another for processing.

		Load balancing: Channels can be used to load-balance work across multiple goroutines. For example, you can use a channel to distribute tasks to multiple worker goroutines, and to collect the results of their work.

		Synchronization: Channels can be used to synchronize the execution of multiple goroutines. For example, you can use a channel to ensure that one goroutine doesn't proceed until another goroutine has completed its work.

		Overall, channels are a powerful tool for building concurrent programs in Go. They provide a simple and safe way for goroutines to communicate and synchronize with each other, and they enable developers to take full advantage of the benefits of concurrent programming.

		Channels in Go provide a safe way to synchronize and communicate between goroutines, but they are not locks in the traditional sense.

		When a goroutine sends or receives data over a channel, it will block until the other end of the channel is ready to receive or send the data. This means that channels provide a form of synchronization that is safe and free from the race conditions and deadlocks that can occur when using locks directly.

		Channels also provide a way to limit the amount of concurrent activity in a system. For example, you can use a channel with a fixed buffer size to limit the number of goroutines that can be active at any given time, or to limit the number of requests that a service can process simultaneously.

		That being said, channels in Go are not a replacement for locks in all situations. In some cases, you may need to use a traditional lock to protect access to shared resources, or to implement more fine-grained synchronization between goroutines.

		Overall, channels in Go are a powerful tool for building concurrent systems that are safe, efficient, and easy to reason about. They provide a simple and safe way to communicate and synchronize between goroutines, and they can be used to implement a wide range of concurrent patterns and architectures.
	*/
	// Declare a channel of integers
	var ch chan int
	// Initialize the channel using the make function
	ch = make(chan int)

	// Send a value on the channel using the <- operator
	go func() {
		ch <- 42
	}()

	// Receive a value from the channel using the <- operator
	val := <-ch
	fmt.Println("The value received from the channel is", val)

}
