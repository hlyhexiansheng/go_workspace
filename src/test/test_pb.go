package main

import (
	"github.com/golang/protobuf/proto"
	"chatmsg"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func main() {
	testMsg := &chatmsg.Msg{
		MsgId: proto.Int32(10),
		MsgInfo:proto.String("thijng"),
		MsgFrom:proto.String("thijng"),
		Money:[]string{"1fjiosjfiosjfiosfjsdoidfjsaiojfisofjaosifjaosidfjsiodfjio", "1", "1", "1"},
	}

	fmt.Println(testMsg)

	//pbBytes := testMarshalPb(testMsg)
	//
	//jsonBytes := testMarshalJson(testMsg)
	//
	//testUnMarshalPb(pbBytes)
	//
	//testUnMarshalJson(jsonBytes)

	testUmarchalFromJava();

}
func testUmarchalFromJava() {

	byteArray, err := ioutil.ReadFile("/tmp/pbdata")
	if err != nil {
		fmt.Println(err)
	}
	msg := &chatmsg.Msg{}
	err = proto.UnmarshalMerge(byteArray, msg)
	if err != nil {
		fmt.Println("error,unmarshal", err)
	}
	fmt.Println(msg)

}
func testMarshalJson(msg *chatmsg.Msg) []byte {
	jsonData, _ := json.Marshal(msg)
	return jsonData
}
func testMarshalPb(msg *chatmsg.Msg) []byte {
	pbData, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("error,marshal", err)
	}
	return pbData
}

func testUnMarshalPb(bytes []byte) {
	msg := &chatmsg.Msg{}
	err := proto.UnmarshalMerge(bytes, msg)
	if err != nil {
		fmt.Println("error,unmarshal", err)
	}
	fmt.Println(msg)
}

func testUnMarshalJson(bytes []byte) {
	msg := &chatmsg.Msg{}
	err := json.Unmarshal(bytes, msg)
	if err != nil {
		fmt.Println("error,unmarshal", err)
	}
	fmt.Println(msg)
}
