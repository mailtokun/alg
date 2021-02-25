package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	out := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// add to channel
	go func() {
		defer wg.Done()
		rand.Seed(time.Now().Unix())
		for i := 0; i < 5; i++ {
			out <- rand.Intn(10)
		}
		close(out)
	}()
	// read from channel
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
