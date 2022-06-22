// You can edit this code!
// Click here and start typing.
package main

import "fmt"
//import "test"
func main() {
	fmt.Println("Select from the options what u want to use:")
	fmt.Println("1.Trignometric Identities ")
	fmt.Println("2.Taylor Expansion")
	fmt.Println("3.Binomial Coefficient")
	fmt.Println("4.Loan Payment")
	fmt.Println("5.Gaussian")

	
	fmt.Println("Select Any:")
	
	var x int
	fmt.Scanln(&x)
    
	if x==1{
		Trignometricidentity()
	}else if x==2{
		TaylorExpansion()
	}else if x==3{
		Binomial()
	}else if x==4{
		LoanPayment()
	}else if x==5{
		Gaussain()
	}else{ 
		fmt.Println("Wrong Input")}

}
      