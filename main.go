package main

import (
	"fmt"
	"multithreadSorting/multithreadingSort"
	"time"
)

func main() {

	timestamp1 := time.Now()

	goroutines := 8

	sortedDoctors, err := multithreadingSort.SortArray(goroutines)
	if err != nil {
		panic("Something went wrong while sorting!")
	}

	timestamp2 := time.Since(timestamp1)

	for i, d := range sortedDoctors {
		fmt.Printf("Doctor %d:\n", i+1)
		fmt.Printf("Name: %s\n", d.Name)
		fmt.Printf("Phone Number: %s\n", d.PhoneNumber)
		fmt.Printf("Email: %s\n", d.Email)
		fmt.Printf("User Agent: %s\n", d.UserAgent)
		fmt.Printf("Hex Color: %s\n", d.HexColor)
		fmt.Println()
	}

	fmt.Printf("Функция выполнилась за %s", timestamp2)

}
