package main

import (
	"toolkits/sys"
	"bufio"
	"bytes"
	"toolkits/file"
	"fmt"
)

func main()  {
	var bs []byte
	bs, _ = sys.CmdOutBytes("sh", "-c", "ss -s")

	reader := bufio.NewReader(bytes.NewBuffer(bs))

	var line []byte;
	for {
		line, _ = file.ReadLine(reader)
		lineStr := string(line[:len(line)])
		fmt.Println(lineStr)
	}
}


