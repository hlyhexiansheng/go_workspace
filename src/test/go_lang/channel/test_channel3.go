package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go fun_1(c)
	go fun_2(c)
	go fun_3(c)
	for {
		select {

		}
	}

}
func fun_1(c chan int) {
	for {
		time.Sleep(time.Duration(1) * time.Second)
		c <- 1
	}
}
func fun_2(c chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("fun_2", v)
		}
	}
}
func fun_3(c chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("fun_3", v)
		}
	}
}
