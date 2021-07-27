package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	arc := runtime.GOARCH

	fmt.Println(os)
	fmt.Println(arc)
}
