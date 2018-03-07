package main

import (
	"agent/logfile"
	"fmt"
)

func main() {
	//test1()
	//test2();
	//test3()
	//test4()
	test5()

}
func test5() {

	//连续不停读，直到整个文件读完，测试的时候把buffter_size改一下
	tailFile := logfile.NewTailFile2("file_node",
		"/Users/noodles/Documents/go_workspace/src/test/read_file/test_material/test5.txt",
		"/Users/noodles/logs",
		0,
		"^\\[INFO", 100, "appName", "domain", "topic", "logType")
	sum := 0
	for ; ; {
		result := tailFile.ReadMultiLine()

		if len(result) == 0 {
			fmt.Println("finish...", sum)
			break
		}

		fmt.Println("length=", len(result))
		for i, v := range result {
			fmt.Println("line:", i, v)
			sum += 1
		}

		tailFile.Commit()
	}

}

func test4() {
	//测试一次buffter读满的问题，测试的时候把buffter_size改一下
	tailFile := logfile.NewTailFile2("file_node",
		"/Users/noodles/Documents/go_workspace/src/test/read_file/test_material/test4.txt",
		"/Users/noodles/logs",
		0,
		"^\\[INFO", 100, "appName", "domain", "topic", "logType")

	result := tailFile.ReadMultiLine()

	fmt.Println("length=", len(result))

	for i, v := range result {
		fmt.Println("line:", i, v)
	}
}

func test3() {
	//测试正则分割行的问题
	tailFile := logfile.NewTailFile2("file_node",
		"/Users/noodles/Documents/go_workspace/src/test/read_file/test_material/test3.txt",
		"/Users/noodles/logs",
		0,
		"^\\[INFO", 100, "appName", "domain", "topic", "logType")

	result := tailFile.ReadMultiLine()

	fmt.Println("length=", len(result))

	for i, v := range result {
		fmt.Println("line:", i, v)
	}
}
func test2() {
	//测试单行最大长度
	tailFile := logfile.NewTailFile2("file_node",
		"/Users/noodles/Documents/go_workspace/src/test/read_file/test_material/test2.txt",
		"/Users/noodles/logs",
		0,
		"^\\[INFO", 10, "appName", "domain", "topic", "logType")

	result := tailFile.ReadMultiLine()

	fmt.Println("length=", len(result))

	for i, v := range result {
		fmt.Println("line:", i, v)
	}

}
func test1() {
	//普通测试
	tailFile := logfile.NewTailFile2("file_node",
		"/Users/noodles/Documents/go_workspace/src/test/read_file/test_material/test1.txt",
		"/Users/noodles/logs",
		0,
		"^\\[INFO", 10, "appName", "domain", "topic", "logType")

	result := tailFile.ReadMultiLine()

	fmt.Println("length=", len(result))

	for _, v := range result {
		fmt.Println(v)
	}
}
