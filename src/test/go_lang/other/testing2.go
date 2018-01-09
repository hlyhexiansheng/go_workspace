package main

import "fmt"

func main() {
	registerDefaultErrorHandler()
}

func registerDefaultErrorHandler() error {
	m := map[string]func(int, int){
		"401": unauthorized,
		"402": paymentRequired,
	}
	fmt.Println(len(m))
	return nil
}

func unauthorized(i1, i2 int) {

}

func paymentRequired(i1, i2 int) {

}

