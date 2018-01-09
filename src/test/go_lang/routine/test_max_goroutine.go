package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Millisecond * time.Duration(1)).C
	var count int64 = 0;
	for {
		<-t
		count = count + 20
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
		go worker(count)
	}
}

func worker(index int64) {
	for {
		fmt.Println(index)
		time.Sleep(time.Second * time.Duration(1000000))
	}
}


