package main

import (
	"fmt"
	"os"
)

var offset int64 = 0

func main() {

	file, err := os.OpenFile("/tmp/a.log", os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("can't open file")
	}
	read_line(file)

}
func read_line(file *os.File) {
	bytes := make([]byte, 10)
	size, _ := file.ReadAt(bytes, offset)
	for i := 0; i < size; i++ {

	}

}
