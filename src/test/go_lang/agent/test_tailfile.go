package main

import (
	"agent/logfile"
	"fmt"
)

func main() {

	tailFile := logfile.NewTailFile("12121", "/Users/noodles/logs/a.txt", "/Users/noodles/logs", 0, true, 91, 100, "appName", "domain", "topic", "logType", 100)
	for {
		line := tailFile.ReadMultiLine()
		if (len(line) == 0) {
			break
		}
		fmt.Println("size=", len(line))
		for _, l := range line {
			fmt.Println(l)
		}
		tailFile.Commit()
	}

}
