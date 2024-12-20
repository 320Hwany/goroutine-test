package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	start := time.Now()
	waitGroup.Add(10)

	for i := 0; i < 10; i++ {
		go printHelloWithGoroutine()
	}

	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %s\n", elapsed)
}

func printHelloWithGoroutine() {
	fmt.Println("hello world")
	time.Sleep(time.Second)
	waitGroup.Done()
}
