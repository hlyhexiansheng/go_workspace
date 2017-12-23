package client

import (
	"net"
	"time"
	"log"
	"io"
	"encoding/binary"
	"fmt"
	"errors"
	"agent/protocal"
	"github.com/golang/protobuf/proto"
	"bytes"
	"sync"
)

const ReconnectDelta int = 10

type TcpClient struct {
	lock          *sync.Mutex

	hostNames     []string
	clients       []net.Conn
	lastPickIndex int
}

var Client *TcpClient

func NewClient(ips []string) *TcpClient {
	Client = &TcpClient{hostNames:ips, lastPickIndex:0, lock:new(sync.Mutex)}
	Client.reconnect();
	return Client
}

func (this *TcpClient) SendMsg(request *protocal.Request, isOneWayTrip bool) (*protocal.Response, bool) {

	reqPbBytes, err := proto.Marshal(request)
	if err != nil {
		log.Printf("marshal request error,[%s]\n", err)
		return nil, false
	}
	//由于这个client，是会被多个routine并发调用的。
	this.lock.Lock()
	defer this.lock.Unlock()

	respBytes, ok := this.SendBytesMsg(reqPbBytes, isOneWayTrip)

	if ok && !isOneWayTrip {
		response := &protocal.Response{}
		err := proto.Unmarshal(respBytes, response)
		if err != nil {
			log.Printf("unmarshal request error,[%s]\n", err)
			return nil, false
		}
		return response, true
	}
	return nil, ok
}

func (this *TcpClient) SendBytesMsg(data []byte, isOneWayTrip bool) ([]byte, bool) {
	var ok bool = false
	var resp []byte

	size := len(this.clients)
	for i := 0; i < size; i++ {
		this.lastPickIndex++
		picked := this.lastPickIndex % size
		if this.clients[picked] == nil {
			continue
		}
		this.clients[picked].SetDeadline(time.Now().Add(5 * time.Second))

		pack_Header := make([]byte, 10)
		binary.BigEndian.PutUint32(pack_Header[0:4], uint32(len(data)) + 6)
		binary.BigEndian.PutUint16(pack_Header[4:6], uint16(0))
		binary.BigEndian.PutUint16(pack_Header[6:8], uint16(0))
		binary.BigEndian.PutUint16(pack_Header[8:10], uint16(2))
		buffer := bytes.NewBuffer(pack_Header)
		buffer.Write(data)

		_, err := this.clients[picked].Write(buffer.Bytes())
		if err == nil {
			if isOneWayTrip {
				ok = true
				break
			} else {
				resp, err = this.readMsg(this.clients[picked])
				if err == nil {
					ok = true
					break
				} else {
					log.Printf("read msg error.[%s]", err)
				}
			}

		}

		log.Println("send data fail", err, this.clients[picked].RemoteAddr(), len(data))
		this.clients[picked] = nil
	}
	//重连机制：
	//1.initial
	//2.send message not OK.
	//3.per ReconnectDelta messages had sent.
	if !ok || (this.lastPickIndex % ReconnectDelta == 0) {
		this.reconnect()
	}
	return resp, ok
}

func (this *TcpClient) readMsg(conn net.Conn) ([]byte, error) {
	pack_length := make([]byte, 4)
	_, err := io.ReadFull(conn, pack_length)
	if (err != nil) {
		return nil, errors.New(fmt.Sprintf("read head message error,error=[%s]", err))
	}
	headSize := binary.BigEndian.Uint32(pack_length)
	data := make([]byte, headSize)
	_, err2 := io.ReadFull(conn, data)
	if err2 != nil {
		return nil, errors.New(fmt.Sprintf("read message error,error=[%s]", err))
	}
	//忽略到前面6个协议相关的字节
	return data[6:], nil
}

func (this *TcpClient)  reconnect() {
	log.Println("reconnect server")
	for _, client := range this.clients {
		if (client != nil) {
			err := client.Close()
			if err != nil {
				log.Println("close client error.", err)
			}
		}
	}

	this.clients = make([]net.Conn, len(this.hostNames))
	for _, host := range this.hostNames {
		conn, err := net.DialTimeout("tcp", host, 1000 * time.Millisecond)
		if (err != nil) {
			log.Printf("connect error,errorInfo=[%s],host=[%s]", err, host)
			continue
		}
		this.clients = append(this.clients, conn)
	}
}





