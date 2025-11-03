package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b, gcdValue int) int {
	return (a * b) / gcdValue
}

func main() {
	a:= 12
	b:= 18

	gcdValue := gcd(a, b)
	lcmValue := lcm(a, b, gcdValue)

	fmt.Printf("GCD of %d and %d is %d\n", a, b, gcdValue)
	fmt.Printf("LCM of %d and %d is %d\n", a, b, lcmValue)
}
