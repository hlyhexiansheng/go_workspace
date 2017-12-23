package logfile

import (
	"agent/protocal"
	"github.com/golang/protobuf/proto"
	"time"
	"strconv"
	"agent/client"
	"log"
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"agent/g"
)

type FileReaderWorker struct {
	FileMap        map[string]*TailFile
	client         *client.TcpClient
	SnowflakeNode  *snowflake.Node

	posFileManager *PosFileManager
	configManager  *DirConfigManager

	pauseFlag      bool
	pauseSemaphore chan int
	stopSemaphore  chan int
}

func NewAndInitFileReadWorker(_posFileManager *PosFileManager, _configManager *DirConfigManager, _stopSemaphore chan int, _client *client.TcpClient) *FileReaderWorker {

	rand.Seed(time.Now().Unix())
	sfNode, _ := snowflake.NewNode(rand.Int63n(1024))

	readWork := &FileReaderWorker{
		pauseFlag:false,
		pauseSemaphore:make(chan int),
		SnowflakeNode:sfNode,
		posFileManager:_posFileManager,
		configManager:_configManager,
		stopSemaphore:_stopSemaphore,
		client:_client}
	readWork.ResetTailFile()
	return readWork
}

func (this *FileReaderWorker) Start() {
	tick := time.NewTicker(time.Second * time.Duration(5)).C
	for {
		select {
		case <-this.stopSemaphore:
			goto END
		case <-tick:
			{
				this.doReadFile()
				if this.pauseFlag {
					this.pauseSemaphore <- 1
				}
			}
		}

	}
	END:

	close(this.pauseSemaphore)
	time.Sleep(time.Second * time.Duration(5))
	this.stopSemaphore <- STOP_SEMAPHORE
	log.Println("stop FileReaderWorker")
}

func (this *FileReaderWorker) Pause() chan int {
	this.pauseFlag = true
	return this.pauseSemaphore
}

func (this *FileReaderWorker) ResumeRead() {
	this.pauseFlag = false
}

func (this *FileReaderWorker) ResetTailFile() {
	this.FileMap = make(map[string]*TailFile, 0)
	for _, fs := range this.posFileManager.PosInfoList {
		fileNode := fs.FileNode
		fullName := fs.FullName
		belong2Dir := fs.Belong2dir
		offset := fs.Offset
		useDefaultLineStart := this.configManager.GetBoolConfigByWatcherDir(belong2Dir, "useDefaultLineStart", false)
		startPrefix := this.configManager.GetIntConfigByWatcherDir(belong2Dir, "startPrefix", 91)
		maxLineSize := this.configManager.GetIntConfigByWatcherDir(belong2Dir, "maxLineSize", 10000)
		maxOnceReadCount := this.configManager.GetIntConfigByWatcherDir(belong2Dir, "maxOnceReadCount", 700)

		appName := this.configManager.GetStringHeaderByWatcherDir(belong2Dir, "appName", "")
		domain := this.configManager.GetStringHeaderByWatcherDir(belong2Dir, "domain", "")
		topic := this.configManager.GetStringHeaderByWatcherDir(belong2Dir, "topic", appName)
		logType := this.configManager.GetStringHeaderByWatcherDir(belong2Dir, "logType", "code")

		tailFile := NewTailFile(fileNode, fullName, belong2Dir,
			offset, useDefaultLineStart, int8(startPrefix), maxLineSize,
			appName, domain, topic, logType, maxOnceReadCount)

		log.Println("%v", tailFile)
		this.FileMap[tailFile.FileNode] = tailFile
	}
}

func (this *FileReaderWorker) doReadFile() {
	for _, tailFile := range this.FileMap {
		this.readFile(tailFile)
	}
}

func (this *FileReaderWorker) readFile(tailFile *TailFile) {
	readCount := 0
	for ; ; {
		if this.pauseFlag || readCount > tailFile.MaxOnceReadCount {
			break
		}

		if lines := tailFile.ReadMultiLine(); len(lines) == 0 {
			break
		} else {
			if ok := this.sendLogs(lines, tailFile); ok {
				tailFile.Commit()
				readCount = readCount + len(lines)
				this.posFileManager.UpdateOffset(tailFile.FileNode, tailFile.Offset)
			} else {
				log.Println("send log error")
				break
			}
		}
	}
}

func (this *FileReaderWorker) sendLogs(lines []string, tailFile *TailFile) bool {
	request := this.assembleLog(lines, tailFile)
	_, OK := this.client.SendMsg(request, true)
	return OK
}

func (this *FileReaderWorker) assembleLog(lines []string, tailFile *TailFile) *protocal.Request {
	logItems := make([]*protocal.LogBean, 0)
	for _, logLine := range lines {
		bean := &protocal.LogBean{}
		bean.LogType = proto.String(tailFile.LogType)
		bean.CollectTime = proto.String(strconv.FormatInt(time.Now().Unix(), 10))
		bean.HostName = proto.String(GetHostName())
		bean.FileName = proto.String(tailFile.FileName)
		bean.AppName = proto.String(tailFile.AppName)
		bean.Domain = proto.String(tailFile.Domain)
		bean.Ip = proto.String(GetIp())
		bean.Topic = proto.String(tailFile.Topic)
		bean.FileOffset = proto.String("0")
		bean.FileNode = proto.String(tailFile.FileNode)
		bean.SortedId = proto.String(this.SnowflakeNode.Generate().String())
		bean.Body = proto.String(logLine)
		logItems = append(logItems, bean)
	}
	baseInfo := &protocal.BaseInfo{ProtocalVersion:proto.Int32(1), Cmd:proto.Int32(g.CMD_REPORT_LOG), ReqId:proto.Int64(110000)}
	request := &protocal.Request{BaseInfo:baseInfo, Logs:logItems}
	return request
}





