package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go accept(c);
	for {
		time.Sleep(time.Second * 2)
		c <- 1
	}
}
func accept(intChan chan int) {
	for v := range intChan {
		fmt.Println("----", v)
	}
}
