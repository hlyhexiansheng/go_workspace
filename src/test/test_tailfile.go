package main

import (
	"agent/logfile"
	"fmt"
)

func main() {
	tailFile := logfile.NewTailFile("12121", "/Users/noodles/logs/cc.txt", "/Users/noodles/logs", 0, false, 91, 100)
	line := tailFile.ReadMultiLine()
	fmt.Println("size=", len(line))
	for i, l := range line {
		fmt.Println("<<", i, "val=", l, ">>")
	}
}
