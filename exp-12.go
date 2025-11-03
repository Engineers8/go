package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Germany", "y"))         
	fmt.Println(strings.ContainsAny("Grossmany", "9"))       
	fmt.Println(strings.ContainsAny("Germany", "gus"))       
	fmt.Println(strings.Contains("Germany", "ger"))          
	fmt.Println(strings.Contains("Germany", "Ger"))          
	fmt.Println(strings.Count("cheese", "e"))                
	fmt.Println(strings.EqualFold("Cat", "cat"))             
	fmt.Println(strings.EqualFold("India", "Indiana"))       
}
