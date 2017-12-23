package main

import (
	"encoding/json"
	"log"
	"fmt"
	"common/model"
	"agent/g"
	"strings"
	"strconv"
	"toolkits/net/httplib"
	"agent/logfile"
)

func main() {
	//testsss()

	//testsingle();
	//syncBuiltinMetrics()

	testlist()
}
func testlist() {

	list := []*logfile.FileStruct{&logfile.FileStruct{FileNode:"1231212", FullName:"aaa.txt", Belong2dir:"/root", Offset:0}, &logfile.FileStruct{FileNode:"212", FullName:"bb.txt", Belong2dir:"/root", Offset:1}}

	bytes, _ := json.Marshal(list)
	fmt.Println(string(bytes))

	data := make([]*logfile.FileStruct, 0);
	json.Unmarshal(bytes, &data)

	fmt.Println(data)
}
func testsingle() {
	fs := &logfile.FileStruct{FileNode:"1231212", FullName:"aaa.txt", Belong2dir:"/root", Offset:2}

	bytes, _ := json.Marshal(fs)

	var data logfile.FileStruct;
	json.Unmarshal(bytes, &data)

	fmt.Println(string(bytes))
	fmt.Println(fs)
	fmt.Println(&data)
}

func testsss() {
	content := []byte("{\"Checffksum\":\"sd\",\"Timestaffmp\":1,\"Metricds\":[{\"Metric\":\"mm1\",\"Tags\":\"ttt1\"},{\"Metric\":\"mm2\",\"Tags\":\"ttt2\"},{\"Metric\":\"mm3\",\"Tags\":\"ttt3\"}]}");
	var m model.BuiltinMetricResponse;
	err := json.Unmarshal(content, &m);
	if err != nil {
		log.Fatalln("parse config file:", err)
	}
	fmt.Println(m)
}

func syncBuiltinMetrics() {

	var ports = []int64{}
	var paths = []string{}
	var procs = make(map[string]map[int]string)
	var urls = make(map[string]string)

	rspnString, err := httplib.PostJSON("http://127.0.0.1:7479/syncBuiltinMetrics", "");
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	var resp model.BuiltinMetricResponse
	err = json.Unmarshal([]byte(rspnString), &resp);
	if err != nil {
		log.Fatalln("parse config file:", err)
		return
	}

	fmt.Println(resp)

	if !resp.IsUpdate {
		fmt.Println("no need to update")
		return
	}

	for _, metric := range resp.Metrics {

		if metric.Metric == g.URL_CHECK_HEALTH {
			arr := strings.Split(metric.Tags, ",")
			if len(arr) != 2 {
				continue
			}
			url := strings.Split(arr[0], "=")
			if len(url) != 2 {
				continue
			}
			stime := strings.Split(arr[1], "=")
			if len(stime) != 2 {
				continue
			}
			if _, err := strconv.ParseInt(stime[1], 10, 64); err == nil {
				urls[url[1]] = stime[1]
			} else {
				log.Println("metric ParseInt timeout failed:", err)
			}
		}

		if metric.Metric == g.NET_PORT_LISTEN {
			arr := strings.Split(metric.Tags, "=")
			if len(arr) != 2 {
				continue
			}

			if port, err := strconv.ParseInt(arr[1], 10, 64); err == nil {
				ports = append(ports, port)
			} else {
				log.Println("metrics ParseInt failed:", err)
			}

			continue
		}

		if metric.Metric == g.DU_BS {
			arr := strings.Split(metric.Tags, "=")
			if len(arr) != 2 {
				continue
			}

			paths = append(paths, strings.TrimSpace(arr[1]))
			continue
		}

		if metric.Metric == g.PROC_NUM {
			arr := strings.Split(metric.Tags, ",")

			tmpMap := make(map[int]string)

			for i := 0; i < len(arr); i++ {
				if strings.HasPrefix(arr[i], "name=") {
					tmpMap[1] = strings.TrimSpace(arr[i][5:])
				} else if strings.HasPrefix(arr[i], "cmdline=") {
					tmpMap[2] = strings.TrimSpace(arr[i][8:])
				}
			}

			procs[metric.Tags] = tmpMap
		}
	}

	fmt.Println("-url-")
	fmt.Println(urls)
	fmt.Println("-port-")
	fmt.Println(ports)
	fmt.Println("-procs-")
	fmt.Println(procs)
	fmt.Println("-patch-")
	fmt.Println(paths)

	g.SetReportUrls(urls)
	g.SetReportPorts(ports)
	g.SetReportProcs(procs)
	g.SetDuPaths(paths)

}



