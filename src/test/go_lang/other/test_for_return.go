package main

import (
	"time"
	"fmt"
	"log"
)

func main() {
	go startLoop()
	select {}
}

func startLoop()  {
	duration := time.Duration(1) * time.Second
	i := 0;
	for  {
		time.Sleep(duration)

		i++
		if i > 3 && i < 6 {
			log.Panic("parse json fail.:")
		}
		fmt.Println(i)

	}
}


