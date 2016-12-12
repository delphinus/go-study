package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const parallels = 10

func heavyWork(n int) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return fmt.Sprintf("work: %d", n)
}

func worker(num int) (chan string, chan bool) {
	wg := sync.WaitGroup{}
	ch := make(chan string)
	fin := make(chan bool)
	go func() {
		for i := 1; i <= num; i++ {
			wg.Add(1)
			go func(i int) {
				ch <- heavyWork(i)
				wg.Done()
			}(i)
		}
		wg.Wait()
		fin <- false
	}()
	return ch, fin
}

func main() {
	ch, fin := worker(parallels)
	for {
		select {
		case result := <-ch:
			fmt.Println(result)
		case <-fin:
			return
		}
	}
}
