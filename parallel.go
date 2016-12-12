package main

import (
	"fmt"
	"math/rand"
	"time"
)

const parallels = 10

func heavyWork(n int) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return fmt.Sprintf("work: %d", n)
}

func worker(num int) chan string {
	ch := make(chan string)
	for i := 1; i <= num; i++ {
		go func(i int) {
			ch <- heavyWork(i)
		}(i)
	}
	return ch
}

func main() {
	ch := worker(parallels)
	for i := 1; i <= parallels; i++ {
		result := <-ch
		fmt.Println(result)
	}
}
