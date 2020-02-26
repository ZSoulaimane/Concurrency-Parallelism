package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Receiving to Channel", i)
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println()
			fmt.Println("Sending from Channel", <-c)
		}
		close(c) //close channel
	}()

	time.Sleep(time.Second)
}
