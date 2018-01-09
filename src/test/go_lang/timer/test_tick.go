package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second * time.Duration(1)).C
	for {
		<-t
		fmt.Println("--------split line ------------")
	}
}
