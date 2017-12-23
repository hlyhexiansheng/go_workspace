package main

import (
	"time"
	"fmt"
	"log"
)

func main() {
	c := make(chan int)
	go fun1(c)
	go fun2(c)
	for {
		select {

		}
	}
}

func fun1(c chan int) {
	for ; ; {
		time.Sleep(time.Duration(5) * time.Second)
		c <- 1
		fmt.Println("pushed event")
	}

}

func fun2(c chan int) {
	tick := time.NewTicker(time.Second * time.Duration(1)).C
	i := 0
	for {
		select {
		case <-tick:
			{
				fmt.Println("tick event")
				i++
			}
		case v1 := <-c:
			fmt.Println("custom event", v1)
			if i > 18 {
				goto END
			}
		}

	}

	END:
	log.Println("stop...")
}
