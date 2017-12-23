package main

import (
	"toolkits/net/httplib"
	"fmt"
)

func main() {
	resopone, error := httplib.PostJSON("http://192.168.56.104:7479/OSMonitor", "");

	fmt.Println("error", error)

	fmt.Println(string(resopone))
}
