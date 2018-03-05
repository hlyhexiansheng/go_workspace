package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
	"log"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.4.4.204:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	_, err = cli.Put(context.TODO(), "test2", "v1")
	_, err = cli.Put(context.TODO(), "test2", "v2")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.Get(context.TODO(), "test2")

	for keyVal, v := range resp.Kvs {
		fmt.Println("---", string(resp.Kvs[keyVal].Value), string(v.Value))
	}
	fmt.Println(resp.Kvs)

	fmt.Println(resp)
}

