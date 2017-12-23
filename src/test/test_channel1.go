package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func1(c)
	go func2(c)

	time.Sleep(time.Duration(10) * time.Second)
}
func func1(c chan int) {
	v := <-c
	fmt.Println(v)
}
func func2(c chan int) {
	time.Sleep(time.Duration(5) * time.Second)
	c <- 1
}
