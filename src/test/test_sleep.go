package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(time.Now().String())

	time.Sleep(time.Duration(8) * time.Second)

	fmt.Println(time.Now().String())
}

