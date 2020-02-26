package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func show(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("fun show:", i)
	}
	defer w.Done()
}
func display(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("fun display:", i)
	}
	defer w.Done()
}
func main() {
	fmt.Println("Calling with sync WaitGroup")
	w.Add(2)
	go show(50)
	go display(50)
	w.Wait()
	for i := 1; i <= 10; i++ {
		fmt.Println("For", i)
	}
}
