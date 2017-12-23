package cron

import (
	"time"
	"agent/funcs"
	"common/model"
	"agent/g"
	"agent/protocal"
	"github.com/golang/protobuf/proto"
	"agent/client"
	"fmt"
)

func InitDataHistory() {
	for {
		funcs.UpdateCpuStat()
		funcs.UpdateDiskStats()
		time.Sleep(g.COLLECT_INTERVAL)
	}
}

func Collect() {

	if !g.Config().Transfer.Enabled {
		return
	}

	if len(g.Config().Transfer.Addrs) == 0 {
		return
	}

	for _, v := range funcs.Mappers {
		go collect(int64(v.Interval), v.Fs)
	}
}

func collect(sec int64, fns []func() []*model.MetricValue) {
	t := time.NewTicker(time.Second * time.Duration(sec)).C
	for {
		<-t

		hostname, err := g.Hostname()
		if err != nil {
			continue
		}

		mvs := []*model.MetricValue{}
		ignoreMetrics := g.Config().IgnoreMetrics

		for _, fn := range fns {
			items := fn()
			if items == nil {
				continue
			}

			if len(items) == 0 {
				continue
			}

			for _, mv := range items {
				if b, ok := ignoreMetrics[mv.Metric]; ok && b {
					continue
				} else {
					mvs = append(mvs, mv)
				}
			}
		}

		now := time.Now().Unix()
		for j := 0; j < len(mvs); j++ {
			mvs[j].Step = sec
			mvs[j].Endpoint = hostname
			mvs[j].Timestamp = now
		}
		request := assembleMetrics(mvs)
		client.Client.SendMsg(request, true)
	}

}

func assembleMetrics(mvs []*model.MetricValue) *protocal.Request {
	logItems := make([]*protocal.Metrics, 0)
	for _, metricsValue := range mvs {
		item := &protocal.Metrics{}
		item.Endpoint = &metricsValue.Endpoint
		item.Metric = &metricsValue.Metric
		item.Step = &metricsValue.Step
		item.Tags = &metricsValue.Tags
		item.Timestamp = &metricsValue.Timestamp
		item.Value = proto.String(InterfaceToString(metricsValue.Value))
		item.Type = &metricsValue.Type
		logItems = append(logItems, item)
	}
	baseInfo := &protocal.BaseInfo{ProtocalVersion:proto.Int32(1), Cmd:proto.Int32(g.CMD_REPORT_METRICS), ReqId:proto.Int64(110000)}
	request := &protocal.Request{BaseInfo:baseInfo, MertricsValue:logItems}
	return request
}

func InterfaceToString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}