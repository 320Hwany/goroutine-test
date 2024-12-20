package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		printHello()
	}

	elapsed := time.Since(start)
	fmt.Printf("걸린 시간: %s\n", elapsed)
}

func printHello() {
	fmt.Println("hello world")
	time.Sleep(time.Second)
}
