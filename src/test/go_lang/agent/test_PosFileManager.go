package main

import (
	"agent/logfile"
	"agent/g"
	"fmt"
)

func main() {
	configFile := "/Users/noodles/Documents/go_workspace/src/agent/cfg.json"
	g.ParseConfig(configFile)

	configManager := &logfile.DirConfigManager{}
	configManager.Init(g.Config().CollectDirs.Dirs)

	posFileManager := &logfile.PosFileManager{FirstLoadSemaphore:true}
	posFileManager.Init(configManager, "/Users/noodles/logs/position.json")
	list := posFileManager.FindDirAllFile(configManager)
	fmt.Println(list)
}
