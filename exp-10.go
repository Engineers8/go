package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	uniqueMap := make(map[int]bool)
	var uniqueArr []int

	for _, val := range arr {
		if !uniqueMap[val] { // if not already present
			uniqueMap[val] = true
			uniqueArr = append(uniqueArr, val)
		}
	}

	fmt.Println("Array after removing duplicates:", uniqueArr)
}
