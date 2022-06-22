package main

import (
	"fmt"
	"math"
)



func TaylorExpansion() {
	var x float64 = 2.0
	var num float64 = 10.0
	var fact float64 = 1.0
	var sum float64 = 0.0
	var i float64
	
	fmt.Println("Enter value of x: ");
	fmt.Scanln(&x);
	fmt.Println("Enter number of terms in taylor series: ");
	fmt.Scanln(&num);

	for i = 1.0; i < num; i++ {
		fact = fact * i
		sum = sum + (math.Pow(x, i) / fact)
	}

	// else{
	//     fmt.Println("Number of terms should be positive")
	// }

	fmt.Println("The value of Tayloe Expansion is:",sum)
}
func main(){
	TaylorExpansion()
}