package main

import (
	"fmt"
	"sync"
)

var waitGroup2 sync.WaitGroup
var money2 = 0
var lock = sync.Mutex{}

func main() {
	waitGroup2.Add(1000)

	for i := 0; i < 1000; i++ {
		go depositWithLock(1000)
	}

	waitGroup2.Wait()
	fmt.Println(money2)
}

func depositWithLock(depositAmount int) {
	lock.Lock()
	defer lock.Unlock()

	money2 += depositAmount
	waitGroup2.Done()
}
