package main

import "fmt"

func main() {
	for {
		fmt.Print("Account balance: ")
		var accountBalance float64
		fmt.Scanln(&accountBalance)
		fmt.Print("Account risk: ")
		var risk float64
		fmt.Scanln(&risk)
		riskInUsd := accountBalance * (risk / 100)
		fmt.Print("Risk: ")
		var x float64
		fmt.Scanln(&x)
		posusd := riskInUsd / (x / 100)
		fmt.Print("price: ")
		var price float64
		fmt.Scanln(&price)
		fmt.Println("\n   position size (usd): ", posusd)
		fmt.Println("\n   position size: ", posusd/price, "\n ")

	}
}
