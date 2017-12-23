package logfile

import (
	"log"
	"time"
	"agent/client"
	"agent/g"
)

var STOP_SEMAPHORE int = 1

type TailDirSource struct {
	Dirs                          []g.WatcherDirConfigDef
	positionFileName              string

	configManager                 *DirConfigManager
	posFileManager                *PosFileManager
	fileReaderWorker              *FileReaderWorker
	client                        *client.TcpClient

	stopWatchDirTaskSemaphore     chan int
	stopSyncPostFileTaskSemaphore chan int
	stopReadWorkerSemaphore       chan int
}

func NewTailDirSource(dirs []g.WatcherDirConfigDef, _positionFileName string, client *client.TcpClient) *TailDirSource {

	tailDirSource := &TailDirSource{}
	tailDirSource.Dirs = dirs
	tailDirSource.positionFileName = _positionFileName
	tailDirSource.client = client

	tailDirSource.stopSyncPostFileTaskSemaphore = make(chan int)
	tailDirSource.stopWatchDirTaskSemaphore = make(chan int)
	tailDirSource.stopReadWorkerSemaphore = make(chan int)
	return tailDirSource
}

func (this *TailDirSource) Init() {
	this.initConfig()
	this.initPosFileManager()
	this.initReaderWorker()
}

func (this *TailDirSource) Start() {
	go this.fileReaderWorker.Start()
	go this.reloadWatchDirTask()
	go this.syncPositionFileTask();
}

func (this *TailDirSource) Stop() {
	log.Println("trying to stop taidir source")
	//send semaphore to each routine,make them stop safely
	this.stopReadWorkerSemaphore <- STOP_SEMAPHORE
	this.stopWatchDirTaskSemaphore <- STOP_SEMAPHORE
	this.stopSyncPostFileTaskSemaphore <- STOP_SEMAPHORE


	//block util routine finish stop.
	<-this.stopWatchDirTaskSemaphore
	<-this.stopSyncPostFileTaskSemaphore
	<-this.stopReadWorkerSemaphore

	close(this.stopWatchDirTaskSemaphore)
	close(this.stopSyncPostFileTaskSemaphore)
	close(this.stopReadWorkerSemaphore)

	//sync position file to disk
	this.posFileManager.SyncDisk()

}

func (this *TailDirSource) reloadWatchDirTask() {
	tick := time.NewTicker(time.Second * time.Duration(100)).C
	for {
		select {
		case <-this.stopWatchDirTaskSemaphore:
			goto END
		case <-tick:
			{
				log.Println("reloadWatchDirTask")

				//1.pause the read routine
				stopNotifyC := this.fileReaderWorker.Pause()
				<-stopNotifyC //let current routine block util read worker routine paused completely finish.

				//2.synchronize file position
				this.posFileManager.SyncDisk()

				//3.reload position file.
				this.posFileManager.Refresh()

				//4.reset the tailFiles
				this.fileReaderWorker.ResetTailFile()

				//5.resume read routine
				this.fileReaderWorker.ResumeRead()
			}
		}
	}
	END:
	log.Println("stop reloadWatchDirTask")
	this.stopWatchDirTaskSemaphore <- STOP_SEMAPHORE
}

func (this *TailDirSource) syncPositionFileTask() {
	tick := time.NewTicker(time.Second * time.Duration(20)).C
	for {
		select {
		case <-this.stopSyncPostFileTaskSemaphore:
			goto END
		case <-tick:
			{
				this.posFileManager.SyncDisk()
			}
		}
	}
	END:
	log.Println("stop syncPositionFileTask")
	this.stopSyncPostFileTaskSemaphore <- STOP_SEMAPHORE
}

func (this *TailDirSource) initReaderWorker() {
	this.fileReaderWorker = NewAndInitFileReadWorker(this.posFileManager, this.configManager, this.stopReadWorkerSemaphore, this.client)
}

func (this *TailDirSource) initPosFileManager() {
	this.posFileManager = &PosFileManager{FirstLoadSemaphore:true}
	this.posFileManager.Init(this.configManager, this.positionFileName)
	log.Println(this.posFileManager.PosInfoList)
}

func (this *TailDirSource) initConfig() {
	this.configManager = &DirConfigManager{}
	this.configManager.Init(this.Dirs)
}

