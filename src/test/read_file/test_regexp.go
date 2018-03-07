package main

import (
	"regexp"
	"fmt"
)

func main() {

	isMatch, err := regexp.MatchString("", "2f")
	fmt.Println(isMatch, err)
}
