package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup
var cnt int

func increment(s string) {
	for i := 1; i <= 10; i++ {
		x := cnt
		x++
		cnt = x
		fmt.Println(s, i, "Counter:", cnt)
	}
	defer w.Done()
}
func main() {
	fmt.Println("Starting")
	w.Add(2)
	go increment("show:") /*Both goroutines accessing increment( ) */
	go increment("display:")
	w.Wait()
	fmt.Println("Final Counter:", cnt)
}
