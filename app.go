package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	// get the relevant data
	revenue, expenses, taxRate = getTheNumbers()
	ebt := revenue - expenses
	profit := ebt - (ebt * taxRate / 100)
	ratio := ebt / profit
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

// getting all the input data from the user
func getTheNumbers() (revenue float64, expenses float64, taxRate float64) {
	fmt.Print("Enter your company's revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your company's expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)
	return
}
