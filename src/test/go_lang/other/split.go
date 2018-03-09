package main

import (
	"regexp"
	"fmt"
)

func main() {

	r := regexp.MustCompile("^\\[\\w+\n")

	s := "[a1\n[a2\n[a3\n"

	result := r.FindAllString(s, -1)
	fmt.Println(len(result))
	for _, v := range result {
		fmt.Println(v)
	}
}
