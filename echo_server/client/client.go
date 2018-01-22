package main

import (
	"net"
	"sync"
	"time"
	"fmt"
)

var (
	svr = "localhost:9999"
	data = []byte("abcdefghijklmnopqrstuvwxyz")
)

var count int64 = 0
var startTime int64 = time.Now().UnixNano()
var mutex sync.Mutex
func addCount() {
	mutex.Lock()
	defer mutex.Unlock()

	count ++
	now := time.Now().UnixNano()
	if count % 1000 == 0 {
		fmt.Printf("speed:%f\n", float64(count) / (float64(now - startTime) / 1e9))
	}
}

func test(conn net.Conn, buf []byte) {
	_, err := conn.Write(data)
	if err != nil {
		panic(err.Error())
	}
	recvLen, err := conn.Read(buf)
	if err != nil {
		panic(err.Error())
	}
	if recvLen != len(data) {
		panic("echo error")
	}
}

func loopTest() {
	conn, err := net.Dial("tcp", svr)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	buf := make([]byte, 100)
	for {
		test(conn, buf)
		addCount()
	}
}

func main() {
	for i := 0; i < 6; i ++ {
		go loopTest()
	}
	select {}
}
