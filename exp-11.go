package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Enter number of integers: ")
	fmt.Scan(&n)

	intArr := make([]int, n)
	fmt.Println("Enter the integers:")
	for i := 0; i < n; i++ {
		fmt.Scan(&intArr[i])
	}

	fmt.Println("\nOriginal integer array:", intArr)

	sort.Sort(sort.Reverse(sort.IntSlice(intArr)))
	fmt.Println("Sorted descending:", intArr)

	var m int
	fmt.Print("\nEnter number of strings: ")
	fmt.Scan(&m)

	strArr := make([]string, m)
	fmt.Println("Enter the strings:")
	for i := 0; i < m; i++ {
		fmt.Scan(&strArr[i])
	}

	fmt.Println("\nOriginal string array:", strArr)

	
	sort.Sort(sort.Reverse(sort.StringSlice(strArr)))
	fmt.Println("Sorted descending:", strArr)
}
