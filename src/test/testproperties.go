package main

import (
	"agent/logfile"
	"fmt"
)

func main() {

	//testLogPropFromFile();

	testLoadFromString();


}
func testLogPropFromFile() {

	pro, _ := logfile.Load_From_File("/Users/noodles/Documents/code/Documents/globalgrow/infrastructure/elog/trunk/flume-parent/flume-plugin/src/test/resources/elog-taildirsource.conf")
	s := pro.String("localproject.sources.tailDir.watchDirs.dir1.domain", "fuck");
	fmt.Println(s)

	//agentName := logfile.GetAgentName(pro)
	//fmt.Println(agentName)
}
func testLoadFromString() {
	con, _ := logfile.ReadFullFileContent("/Users/noodles/Documents/code/Documents/globalgrow/infrastructure/elog/trunk/flume-parent/flume-plugin/src/test/resources/elog-taildirsource.conf")
	pro, _ := logfile.Load_From_String(con)

	s := pro.String("localproject.sources.tailDir.watchDirs.dir1.domain", "fuck");
	fmt.Println(s)

	//	agentName := logfile.GetAgentName(pro)
	//	fmt.Println(agentName)
}
