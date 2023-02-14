package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current date and time is ", currentTime)
	// Current Date
	y, m, d := currentTime.Date()
	fmt.Println("Current date ", y, m, d)

	// print date in format
	// Require to refer go date and time layout
	// Reference: https://go.dev/src/time/format.go
	fmt.Println("Current date by using format", currentTime.Format("2006-01-02"))

	// Sprintf method
	currentYear := fmt.Sprintf("%d", currentTime.Year())
	fmt.Println("Current year", currentYear)
}
