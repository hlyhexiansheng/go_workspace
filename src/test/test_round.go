package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	for i := 0; i < 1000000; i++ {
		data := rand.Int63n(1024)
		fmt.Println(data)
	}

}
