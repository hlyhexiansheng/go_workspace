package main

import (
	"time"
	"fmt"
)

func main() {
	c := make(chan int)
	go ff1(c)
	go ff2(c)
	for {
		select {

		}
	}
}

func ff1(c chan int) {
	for {
		time.Sleep(time.Duration(1) * time.Second)
		c <- 1
		v := <-c
		fmt.Println("ff1 recieved", v)
	}
}
func ff2(c chan int) {
	for {
		time.Sleep(time.Duration(2) * time.Second)
		v := <-c
		fmt.Println("ff2 recieved", v)
		c <- 1
	}
}


