package main

import "fmt"

//Creating a Profit Calculator in Go as a project to learn its basics
func main() {
	println("Welcome to the Profit Calculator!")
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
	EAT := EBT - (EBT * (taxRate/100))

	fmt.Println("Earnings Before Tax: ", EBT)
	fmt.Println("Earnings After Tax: ", EAT)
}
