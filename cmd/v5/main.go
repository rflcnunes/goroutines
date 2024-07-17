package main

import (
	"fmt"
	"time"
)

// Channel buffer
func main() {
	channel := make(chan int, 100)
	go setList(channel)

	for v := range channel {
		fmt.Println("Receive " + fmt.Sprint(v))
		time.Sleep(1 * time.Second)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("Send " + fmt.Sprint(i))
	}

	close(channel)
}
