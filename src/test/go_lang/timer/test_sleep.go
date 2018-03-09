package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(time.Now().String())

	now := time.Now()

	time.Sleep(time.Duration(1) * time.Second)

	fmt.Println(time.Now().String())

	fmt.Println((time.Now().UnixNano() - now.UnixNano()) / 1000 / 1000)

	s := ""
	fmt.Println(s[0])
}

