package main

import (
	"fmt"
	"time"
	"toolkits/net/httplib"
	"io/ioutil"
	"strings"
)

type TestInfo struct {
	status   bool
	duration int64
}

var routineCount = 2
var requestTime = 20

func main() {

	responseChan := make(chan TestInfo)
	finishChan := make(chan int)

	rspnCount := 0
	rspnFailCount := 0

	finishCount := 0

	for i := 0; i < routineCount; i++ {
		go requestLoop(responseChan, finishChan)
	}

	for {
		select {
		case t := <-responseChan:
			{
				rspnCount++

				if t.status == false {
					fmt.Println("rspn=", rspnCount, ",val=", t)
					rspnFailCount++
				}
			}
		case <-finishChan:
			{
				finishCount++
				if finishCount == routineCount {
					goto LABEL_END
				}
			}
		}
	}
	LABEL_END:
	fmt.Println("finish all")
	fmt.Println("rspnCount=", rspnCount, ",failCount=", rspnFailCount)

}

func requestLoop(responseChan chan TestInfo, finishChan chan int) {
	t := time.NewTicker(time.Second * time.Duration(requestTime)).C
	for {
		select {
		case <-t:
			{
				finishChan <- 1
				return
			}
		default:
			{
				doRequestUrl(responseChan)
			}
		}

	}
}

func doRequestUrl(c chan TestInfo) {
	startTime := getUnixTime()

	request := httplib.Post("http://10.4.4.204:8100")
	request.Param("app_key", "gateway");
	request.Param("method", "com.order.query.query_order");
	request.Param("module", "debug_upstream");
	request.Param("timestamp", "2018-01-04 08:47:45");
	request.Param("sign", "03177f66978703d7c11163c88f882729");

	rspn, err := request.Response();

	endTime := getUnixTime()

	if err != nil {
		fmt.Println("request error.", err)
		c <- TestInfo{status:false, duration:endTime - startTime}
	} else {
		bytes, _ := ioutil.ReadAll(rspn.Body)

		isOk := strings.Contains(string(bytes[:]), "business_upstream")
		if isOk == false {
			fmt.Println("request:--", string(bytes[:]))
		}
		c <- TestInfo{status:isOk, duration:endTime - startTime}
	}

}

func getUnixTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

