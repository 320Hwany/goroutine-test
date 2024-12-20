package main

import (
	"fmt"
	"sync"
)

var waitGroup1 sync.WaitGroup
var money1 = 10000

func main() {
	waitGroup1.Add(2)
	go deposit(2000)
	go withdraw(12000)
	waitGroup1.Wait()

	fmt.Println(money1)
}

func deposit(depositAmount int) {
	money1 += depositAmount
	waitGroup1.Done()
}

func withdraw(withdrawAmount int) {
	if money1 < withdrawAmount {
		fmt.Println("잔액이 부족합니다.")
	} else {
		money1 -= withdrawAmount
	}

	waitGroup1.Done()
}
