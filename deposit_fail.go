package main

import (
	"fmt"
	"sync"
)

var waitGroup1 sync.WaitGroup
var money1 = 0

func main() {
	waitGroup1.Add(1000)

	for i := 0; i < 1000; i++ {
		go deposit(1000)
	}

	waitGroup1.Wait()
	fmt.Println(money1)
}

func deposit(depositAmount int) {
	money1 += depositAmount
	waitGroup1.Done()
}
