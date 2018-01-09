package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	testUnixTimeStamp()
}
func testUnixTimeStamp() {

	fmt.Println(strconv.FormatInt(time.Now().Unix(),10))
}
