package logfile

import (
	"strconv"
	"agent/g"
)

type DirConfigManager struct {
	WatDirConfDef map[string]*g.WatcherDirConfigDef
}

func (this *DirConfigManager) Init(dirs []g.WatcherDirConfigDef) {
	this.WatDirConfDef = make(map[string]*g.WatcherDirConfigDef)
	for index, _ := range dirs {
		this.WatDirConfDef[dirs[index].Path] = &dirs[index]
	}
	return
}

func (this *DirConfigManager) GetStringConfigByWatcherDir(watchDir string, configKey string, defaultVal string) (result string) {
	result = defaultVal
	for path, configDef := range this.WatDirConfDef {
		if (path == watchDir) {
			if result = configDef.Config[configKey]; result == "" {
				result = defaultVal
			}
			return
		}
	}
	return
}

func (this *DirConfigManager) GetBoolConfigByWatcherDir(watchDir string, configKey string, defaultVal bool) (bool) {
	sr := this.GetStringConfigByWatcherDir(watchDir, configKey, strconv.FormatBool(defaultVal))
	result, err := strconv.ParseBool(sr)
	if (err != nil) {
		result = defaultVal
	}
	return result
}

func (this *DirConfigManager) GetHeadersByWatchDir(watchDir string) map[string]string {
	for path, configDef := range this.WatDirConfDef {
		if (path == watchDir) {
			return configDef.Header
		}
	}
	return nil
}

func (this *DirConfigManager) GetStringHeaderByWatcherDir(watchDir string, headerKey string, defaultVal string) (result string) {
	result = defaultVal
	for path, configDef := range this.WatDirConfDef {
		if (path == watchDir) {
			if result = configDef.Header[headerKey]; result == "" {
				result = defaultVal
			}
			return
		}
	}
	return
}

func (this *DirConfigManager) GetIntConfigByWatcherDir(watchDir string, configKey string, defaultVal int) (int) {
	sr := this.GetStringConfigByWatcherDir(watchDir, configKey, strconv.FormatInt(int64(defaultVal), 10))
	result, err := strconv.ParseInt(sr, 10, 32)
	if (err != nil) {
		result = int64(defaultVal)
	}
	return int(result)
}
