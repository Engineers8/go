package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of rows: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := i; j < n; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		for j := i - 1; j >= 1; j-- {
			fmt.Print(j)
		}

		fmt.Println()
	}
}
