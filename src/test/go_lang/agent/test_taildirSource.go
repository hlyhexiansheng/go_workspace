package main

import (
	"agent/logfile"
	"log"
	"fmt"
	"agent/client"
	"agent/g"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[")

	testInit()
}
func testInit() {

	client := client.NewClient([]string{"localhost:4142"})
	positionFileName := "/Users/noodles/logs/position.json"

	configFile := "/Users/noodles/Documents/go_workspace/src/agent/cfg.json"
	g.ParseConfig(configFile)

	source := logfile.NewTailDirSource(g.Config().CollectDirs.Dirs, positionFileName, client)
	source.Init()
	source.Start()

	for {
		s := ""
		fmt.Scanf("%s", &s)
		fmt.Println(s)
		if s == "stop" {
			source.Stop()
		}
	}
}
