package main

import (
	"fmt"
	"flag"
	"os"
	"agent/g"
	"agent/funcs"
	"agent/cron"
	"log"
	"agent/logfile"
	"agent/client"
	"toolkits/file"
)

func main() {

	log.Println("start elog-falcon")

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	g.InitRootDir()
	g.InitLocalIps()

	funcs.BuildMappers()

	go cron.InitDataHistory()

	cron.SyncBuiltinMetrics()
	cron.Collect()

	tcpClient := client.NewClient(g.Config().Transfer.Addrs)

	source := logfile.NewTailDirSource(g.Config().CollectDirs.Dirs, file.SelfDir() + "/postion.json", tcpClient)
	source.Init()
	source.Start()

	select {}
}
