package main

import (
	"agent/logfile"
	"fmt"
	"time"
	"os"
	"strings"
)

func main() {
	//test1()
	//test2();
	test3()
	//test4()
	//test5()
	//test6()
	//test7();
}

func test7() {
	//测试单纯读文件速度
	file, _ := os.OpenFile("/Users/noodles/logs/a.txt", os.O_RDONLY, 0444)
	var offset_index int64 = 0
	bytes := make([]byte, 1024 * 1024)
	startTime := time.Now()
	for {
		size, _ := file.ReadAt(bytes, offset_index)
		if (size == 0) {
			break
		}

		result := strings.Split(string(bytes[:size]), "\n")
		for _, v := range result {
			fmt.Println(v)
		}
		offset_index += int64(size)

	}
	fmt.Println(time.Now().Unix() - startTime.Unix())

}

func test6() {

	//测试读大文件
	tailFile := logfile.NewTailFile2("file_node",
		"/Users/noodles/logs/a.txt",
		"/Users/noodles/logs",
		0,
		1,
		"[", 1000, 1000, "appName", "domain", "topic", "logType")
	sum := 0
	startTime := time.Now()

	for ; ; {
		result, _ := tailFile.ReadMultiLine()

		if len(result) == 0 {
			fmt.Println("finish...", sum)
			break
		}

		//for i, v := range result {
		//	fmt.Println("<<<-----line:", i, v , ">>>>>>>>>>")
		//	sum += 1
		//}

		tailFile.Commit()
	}

	fmt.Println(time.Now().Unix() - startTime.Unix())


}

func test5() {

	//连续不停读，直到整个文件读完，测试的时候把buffter_size改一下
	tailFile := logfile.NewTailFile2("file_node",
		"/Users/noodles/Documents/go_workspace/src/test/read_file/test_material/test5.txt",
		"/Users/noodles/logs",
		0,
		1,
		"^\\[INFO", 100, 1000, "appName", "domain", "topic", "logType")
	sum := 0
	for ; ; {
		result, _ := tailFile.ReadMultiLine()

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
		0,
		"^\\[INFO", 100, 1000, "appName", "domain", "topic", "logType")

	result, _ := tailFile.ReadMultiLine()

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
		1,
		"[INFO", 100, 1000, "appName", "domain", "topic", "logType")

	result, _ := tailFile.ReadMultiLine()

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
		0,
		"^\\[INFO", 10, 1000, "appName", "domain", "topic", "logType")

	result, _ := tailFile.ReadMultiLine()

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
		0,
		"^\\[INFO", 10, 1000, "appName", "domain", "topic", "logType")

	result, offset := tailFile.ReadMultiLine()

	fmt.Println("length=", len(result))

	for i, v := range result {
		fmt.Println(v, "offset", offset[i])
	}
}
