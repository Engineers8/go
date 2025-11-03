package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)

	var numbers = make([]int, n)
	var sum int

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
		sum += numbers[i]
	}

	average := sum / n 

	fmt.Println("The average is:", average)
}
