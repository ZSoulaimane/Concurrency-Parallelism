package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup
var cnt int
var mutex sync.Mutex //create mutex type variable
func increment(s string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(3 * time.Millisecond)
		mutex.Lock() //Locking while incrementing
		x := cnt
		x++
		cnt = x
		fmt.Println(s, i, "Counter:", cnt)
		mutex.Unlock() //Unlocking
	}
	defer w.Done()
}
func main() {
	fmt.Println("Starting")
	w.Add(2)
	go increment("show:")
	go increment("display:")
	w.Wait()
	fmt.Println("Final Counter:", cnt)
}
