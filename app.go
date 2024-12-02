package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	// get the relevant data
	fmt.Print("Enter your company's revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your company's expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	profit := ebt - (ebt * taxRate / 100)
	ratio := ebt / profit
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
