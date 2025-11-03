package main

import (
	"fmt"
	"math"
)

func main() {
	var num [10]float64
	var sum, mean, variance, sd float64

	fmt.Println("**** Enter 10 Elements *****")
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&num[i])
		sum += num[i]
	}

	mean = sum / 10

	for i := 0; i < 10; i++ {
		variance += math.Pow(num[i]-mean, 2)
	}
	variance = variance / 10
	sd = math.Sqrt(variance)

	fmt.Println("\nThe Sum is:", sum)
	fmt.Println("The Mean is:", mean)
	fmt.Println("The Variance is:", variance)
	fmt.Println("The Standard Deviation is:", sd)
}
