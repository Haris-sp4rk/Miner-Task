package main

import (
	"fmt"
	"math"
)

func main() {
	Binomial()
}

func factorial(n int) float64 {
	var factVal float64 = 1 // float64 is the set of all unsigned 64-bit integers.

	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= float64(i) // mismatched types int64 and int
		}

	}
	return factVal /* return from function*/
}

func Binomial() {
	fmt.Println()
	fmt.Println("Binomial Coefficient Calculator.")
	fmt.Println("================================")
	fmt.Println()

	var x, y int
	// var y int
	fmt.Println("Enter First Integer: ")
	fmt.Scanln(&x)

	fmt.Println("Enter Second Integer: ")
	fmt.Scanln(&y)

	var bino float64
	bino = math.Floor(factorial(x) / factorial(y) / factorial(x-y))

	fmt.Println("Binomial Coefficient: ", bino)

	// Logic:
	// flor( fac(x) / fac(y) / fac(x - y) )
}
