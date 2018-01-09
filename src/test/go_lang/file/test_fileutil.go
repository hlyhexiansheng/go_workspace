package main

import (
	"agent/logfile"
	"fmt"
)

func main() {
	//testGetLogFile();

	//testFileExist()

	//testIsDir()

	//testIsSymlink()

	//ResolveLink()

	//testSelfPath()

	//testOpenFileAndReadAll();

	testWriteFile()

}
func ResolveLink() {
	dir, _ := logfile.ResolveLink("/Users/noodles/logslink")
	fmt.Println(dir)
}
func testIsSymlink() {

	var is = logfile.IsSymlink("/Users/noodles/logs")
	fmt.Println(is)
}
func testIsDir() {
	var isdir = logfile.IsDir("/Users/noodles")
	fmt.Println(isdir)
}
func testFileExist() {
	var exixt bool = logfile.FileIsExist("/Users/noodles/logslinkfsda")
	fmt.Println(exixt)
}

func testGetLogFile() {
	filelist := logfile.GetAllFile("/Users/noodles/logs", "txt$")
	for i := 0; i < len(filelist); i++ {
		fmt.Println(filelist[i].FullName)
	}

	for ix, season := range filelist {
		fmt.Printf("Season %d is: %s, %s\n", ix, season.FullName, season.Belong2dir)
	}

	fmt.Println(filelist)
}

func testSelfPath() {
	fmt.Println(logfile.SelfPath())
}

func testOpenFileAndReadAll() {
	content, _ := logfile.OpenFileAndReadAll("/Users/noodles/logs/0000.txt")
	fmt.Println(content)
}

func testWriteFile() {
	logfile.OpenFileAndWriteAll("/Users/noodles/logs/0000.txt", "23")
}