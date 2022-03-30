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
		fmt.Println("\n   ", riskInUsd/(x/100))

	}
}
