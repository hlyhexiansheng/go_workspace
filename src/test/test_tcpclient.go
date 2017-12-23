package main

import (
	"fmt"
	"agent/client"
	"log"
	"agent/protocal"
	"github.com/golang/protobuf/proto"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[")
	//test1()
	test2()
}
func test2() {

	client := client.NewClient([]string{"localhost:4142"})

	baseInfo := &protocal.BaseInfo{ProtocalVersion:proto.Int32(1), Cmd:proto.Int32(2), ReqId:proto.Int64(110000)}

	logs := []*protocal.LogBean{{FileName:proto.String("hello"), SortedId:proto.String("888899999")}, {FileName:proto.String("hello222"), SortedId:proto.String("323")}}
	request := &protocal.Request{BaseInfo:baseInfo, Logs:logs}

	resp, ok := client.SendMsg(request, true)
	fmt.Println(resp, ok)
}
func test1() {

	client := client.NewClient([]string{"localhost:8887"})
	for ; ; {
		s := ""
		fmt.Scanf("%s\n", &s)
		data, ok := client.SendBytesMsg([]byte(s), false)
		if ok {
			fmt.Println("DATA=", string(data))
		} else {
			fmt.Println(ok)
		}

	}
}