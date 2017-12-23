package main

import (
	"agent/logfile"
	"fmt"
)

func main() {
	testCInit()
}

func testCInit() {
	rawConfigStr, _ := logfile.ReadFullFileContent("/Users/noodles/Documents/code/Documents/globalgrow/infrastructure/elog/trunk/flume-parent/flume-plugin/src/test/resources/elog-taildirsource.conf")

	configManager := &logfile.DirConfigManager{}
	configManager.Init(rawConfigStr)

	//fmt.Println(configManager.WatDirConfDef)
	//fmt.Println(configManager.GetAgentName())
	//fmt.Println(configManager.Prop)
	//fmt.Println(configManager.GetStringConfigByWatcherDir("/Users/noodles/logs2", "useDefaultLineStart","nothing"))
	//fmt.Println(configManager.GetBoolConfigByWatcherDir("/Users/noodles/logs2", "useDefaultLineStart", true))
	fmt.Println(configManager.GetHeadersByWatchDir("/Users/noodles/log") == nil)
}
