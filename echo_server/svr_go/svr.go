package main

import (
	"net"
)

func work(conn net.Conn) {
	buf := make([]byte, 10000)
	for {
		recvLen, err := conn.Read(buf)
		if err != nil {
			panic(err.Error())
		}
		_, err = conn.Write(buf[: recvLen])
		if err != nil {
			panic(err.Error())
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err.Error())
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}
		go work(conn)
	}
}
