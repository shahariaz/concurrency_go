package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int

	fmt.Printf("Initial account balance: $%d", bankBalance)

	incomes := []Income{
		{"Salary", 5000},
		{"Freelance", 2000},
		{"Investment", 1500},
		{"Gift", 800},
		{"Bonus", 1200},
	}

	for i, income := range incomes {

		go func(i int, income Income) {

		}(i, income)
	}

}
