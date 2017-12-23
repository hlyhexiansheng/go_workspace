package g

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"toolkits/file"
	"fmt"
)

type TransferConfig struct {
	Enabled  bool     `json:"enabled"`
	Addrs    []string `json:"addrs"`
	Interval int      `json:"interval"`
	Timeout  int      `json:"timeout"`
}

type HttpConfig struct {
	Enabled  bool   `json:"enabled"`
	Listen   string `json:"listen"`
	Backdoor bool   `json:"backdoor"`
}

type WatcherDirConfigDef struct {
	Path   string `json:"path"`
	Header map[string]string `json:"header"`
	Config map[string]string `json:"config"`
}
func (this *WatcherDirConfigDef) String() string {
	return fmt.Sprintf("<[path:%s], [Header:%s], [Config:%s]>\n", this.Path, this.Header, this.Config, )
}


type CollectorConfig struct {
	IfacePrefix []string `json:"ifacePrefix"`
}

type CollectLogConfig struct {
	Enable bool        `json:"enabled"`
	Dirs   []WatcherDirConfigDef `json:"dirs"`
}

type GlobalConfig struct {
	Debug           bool             `json:"debug"`
	Hostname        string           `json:"hostname"`
	IP              string           `json:"ip"`
	Transfer        *TransferConfig  `json:"transfer"`
	Http            *HttpConfig      `json:"http"`
	Collector       *CollectorConfig `json:"collector"`
	IgnoreMetrics   map[string]bool  `json:"ignore"`
	CollectInterval int `json:"collectInterval"`
	CollectDirs     *CollectLogConfig        `json:"collectDirs"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("ERROR: os.Hostname() fail", err)
	}
	return hostname, err
}

func IP() string {
	ip := Config().IP
	if ip != "" {
		// use ip in configuration
		return ip
	}

	if len(LocalIps) > 0 {
		ip = LocalIps[0]
	}

	return ip
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
