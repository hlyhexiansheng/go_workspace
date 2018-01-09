package main

import "fmt"

func main() {
	var a int = -9999999999999999
	interface2string(a)
	interface2string(1.2112121212312323)
	interface2string("23.12")
	interface2string("false")
	interface2string(true)
	interface2string(nil)
}

func interface2string(v interface{}) {
	s := fmt.Sprintf("%v", v)
	fmt.Println(s)
}

