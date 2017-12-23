package main

import (
	"log"
	"strings"
	"toolkits/file"
	"bytes"
	"bufio"
	"toolkits/sys"
	"fmt"
)

func main() {
	ok,_ := probeUrl("http://www.baidu.com","1");
	fmt.Println(ok)
}

func probeUrl(furl string, timeout string) (bool, error) {
	bs, err := sys.CmdOutBytes("curl", "--max-filesize", "102400", "-I", "-m", timeout, "-o", "/dev/null", "-s", "-w", "%{http_code}", furl)
	if err != nil {
		log.Printf("probe url [%v] failed.the err is: [%v]\n", furl, err)
		return false, err
	}
	reader := bufio.NewReader(bytes.NewBuffer(bs))
	retcode, err := file.ReadLine(reader)
	if err != nil {
		log.Println("read retcode failed.err is:", err)
		return false, err
	}
	if strings.TrimSpace(string(retcode)) != "200" {
		log.Printf("return code [%v] is not 200.query url is [%v]", string(retcode), furl)
		return false, err
	}
	return true, err
}
