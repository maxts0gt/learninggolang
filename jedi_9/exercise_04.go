package main

import (
	"fmt"
	"runtime"
	"sync"
)

//https://tour.golang.org/concurrency/9
//https://stackoverflow.com/questions/38059618/go-goroutine-lock-and-unlock

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
