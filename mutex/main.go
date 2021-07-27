package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	incrementer := 0

	gs := 100

	wg.Add(100)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)

}
