package main

import (
	"fmt"
	"io"
)

type Interface interface {
	String() string
	Set(s string)
}

type Message struct {
	msg string
}

func (s Message) String() string {
	return "Message: " + s.msg
}

func (s Message) Set(newString string){
	s.msg = newString;

}

func main() {
	m1 := Message{"hello world"}
	var i1 Interface = m1;

	if sv, ok := i1.(io.Reader); ok {
		fmt.Printf("v implements String(): %s\n", sv) // note: sv, not v
	}
}