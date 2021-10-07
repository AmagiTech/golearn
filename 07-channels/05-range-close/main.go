package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
