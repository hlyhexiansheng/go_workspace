package cron

//同步自定义metrics
func SyncBuiltinMetrics() {
	go syncBuiltinMetrics()
}

func syncBuiltinMetrics() {

	//g.SetReportUrls(urls)
	//g.SetReportPorts(ports)
	//g.SetReportProcs(procs)
	//g.SetDuPaths(paths)

}
