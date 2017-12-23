package main

import (
	"agent/protocal"
	"github.com/golang/protobuf/proto"
	"fmt"
	"os"
)

func main() {


	baseInfo := &protocal.BaseInfo{ProtocalVersion:proto.Int32(1), Cmd:proto.Int32(2), ReqId:proto.Int64(110000)}

	logs := []*protocal.LogBean{{FileName:proto.String("hello"), SortedId:proto.String("1111111111111")}, {FileName:proto.String("hello222"), SortedId:proto.String("323")}}
	request := &protocal.Request{BaseInfo:baseInfo, Logs:logs}

	fmt.Println(logs)
	pbBytes, err := proto.Marshal(request)
	if err != nil {
		fmt.Println("....")
	}

	file, err := os.OpenFile("/tmp/request.pb", os.O_CREATE | os.O_RDWR | os.O_SYNC | os.O_TRUNC, 0666)
	file.Write(pbBytes)

	request2 := &protocal.Request{};
	err = proto.Unmarshal(pbBytes, request2)

	fmt.Println(request2)
}
