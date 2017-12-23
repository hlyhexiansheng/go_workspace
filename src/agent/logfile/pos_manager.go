package logfile

import (
	"encoding/json"
	"log"
	"errors"
	"agent/g"
)

type PosFileManager struct {
	FirstLoadSemaphore bool
	PosFileName        string

	PosInfoList        []*FileStruct

	ConfigManager      *DirConfigManager
}

func (this *PosFileManager) Init(configM *DirConfigManager, posFileName string) {
	this.FirstLoadSemaphore = true
	this.ConfigManager = configM
	this.PosFileName = posFileName

	this.Refresh()

	this.FirstLoadSemaphore = false
}

func (this *PosFileManager) Refresh() {
	allList := this.findDirAllFile(this.ConfigManager.WatDirConfDef)

	oldPosInfoList := this.loadOldPosInfo()

	this.mergeResult(allList, oldPosInfoList)

	this.PosInfoList = allList;

	this.SyncDisk()
}

func (this *PosFileManager) UpdateOffset(fileNode string, offset int64) {
	for _, fileStruct := range this.PosInfoList {
		if fileStruct.FileNode == fileNode {
			fileStruct.Offset = offset
		}
	}
}

func (this *PosFileManager) SyncDisk() bool {
	byteArray, err := json.Marshal(this.PosInfoList)
	if err != nil {
		log.Println("SyncDisk error", err)
		return false
	}
	ok, _ := OpenFileAndWriteAll(this.PosFileName, string(byteArray))
	return ok
}

func (this *PosFileManager) mergeResult(allList, posInfoList []*FileStruct) {
	for _, fs := range posInfoList {
		this.setOffset(fs, allList)
	}
}

func (this *PosFileManager) setOffset(oldFs *FileStruct, allList []*FileStruct) {
	for _, f := range allList {
		if f.FileNode == oldFs.FileNode {
			f.Offset = oldFs.Offset
		}

		if this.FirstLoadSemaphore &&
			this.ConfigManager.GetBoolConfigByWatcherDir(f.Belong2dir, "isReadFromEnd", false) {
			f.Offset = GetFileLastOffset(f.FullName)
		}

	}
}

func (this *PosFileManager) loadOldPosInfo() []*FileStruct {
	byteArray, err := OpenFileAndReadAll(this.PosFileName)
	if (err != nil) {
		log.Println("read postion file error.", err)
		panic("read postion file error." )
		return nil
	}
	posList := make([]*FileStruct, 0);
	if len(byteArray) > 0 {
		json.Unmarshal(byteArray, &posList)
	}
	return posList
}

func (this *PosFileManager) findDirAllFile(defMap map[string]*g.WatcherDirConfigDef) []*FileStruct {
	fsl, err := this.getAllFileByWatchConfig(defMap)
	if err != nil {
		log.Panic("parse to prop error")
	}
	return fsl
}

func (this *PosFileManager) getAllFileByWatchConfig(dirConfig map[string]*g.WatcherDirConfigDef) ([]*FileStruct, error) {
	if len(dirConfig) == 0 {
		return nil, errors.New("the dirs size is 0")
	}
	list := make([]*FileStruct, 0)
	for path, config := range dirConfig {
		tmpList := GetAllFile(path, config.Config["filterKey"])
		list = append(list, tmpList...)
	}
	return list, nil
}



