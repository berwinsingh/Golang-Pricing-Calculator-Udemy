package main

import "fmt"

//Creating a Profit Calculator in Go as a project to learn its basics
func main() {
	fmt.Println("Welcome to the Profit Calculator!")

	var revenue float64
	var cost float64
	var taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the cost: ")
	fmt.Scan(&cost)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - cost
	EAT := EBT * (1 - (taxRate / 100))
	ratio := EBT / EAT

	fmt.Println("Earnings Before Tax: ", EBT)
	fmt.Println("Earnings After Tax: ", EAT)
	fmt.Printf("Ratio: %.2f\n", ratio) //Restricting the decimal points to 2
}
