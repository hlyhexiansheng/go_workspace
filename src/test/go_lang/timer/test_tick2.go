package main

import (
	"time"
	"fmt"
)

func main() {
	tick := time.NewTicker(time.Second * time.Duration(1)).C
	for {
		select {
		case v := <-tick:
			{
				fmt.Println("tick....", v, len(tick))
				time.Sleep(time.Duration(5) * time.Second)
			}
		}

	}
}
