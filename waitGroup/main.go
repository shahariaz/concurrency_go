package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printHello(wg *sync.WaitGroup) {
	defer wg.Done()
	defer wg.Done()
	fmt.Println("Hello World")
}

func main() {

	wg.Add(2)
	go printHello(&wg)

	wg.Wait()

}
