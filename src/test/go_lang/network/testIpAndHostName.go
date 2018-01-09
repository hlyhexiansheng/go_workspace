package main

import (
	"toolkits/net"
	"fmt"
)

func main() {
	//testIPGet1()
	testIPGet2()

}
func testIPGet2() {
	ip, err := net.IntranetIP()
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(ip)
}

func testIPGet1() {

	//name, _ := os.Hostname()
	//fmt.Println(name)
	//
	//ifaces, err := net.Interfaces()
	//// handle err
	//for _, i := range ifaces {
	//	addrs, _ := i.Addrs()
	//	for _, addr := range addrs {
	//		var ip net.IP
	//		switch v := addr.(type) {
	//		case *net.IPNet:
	//			ip = v.IP
	//		case *net.IPAddr:
	//			ip = v.IP
	//		}
	//		fmt.Println(ip)
	//	}
	//
	//}
	//if err != nil {
	//	fmt.Println(err)
	//}
}
