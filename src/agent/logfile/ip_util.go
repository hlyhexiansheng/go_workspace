package logfile

import (
	"toolkits/net"
	"log"
	"os"
)

var ip string

func GetIp() string {
	if ip == "" {
		if ips, err := net.IntranetIP(); err != nil || len(ips) == 0 {
			ip = "127.0.0.1"
			log.Println("get Ip error.")
		} else {
			ip = ips[0]
		}
	}
	return ip
}

var hostName string

func GetHostName() string {
	if hostName == "" {
		if host, err := os.Hostname(); err != nil || host == "" {
			host = "localhost"
		} else {
			hostName = host
		}
	}
	return hostName
}
