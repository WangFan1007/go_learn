package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var message = make(chan string)

func WriteMessageToClient(client Client, conn net.Conn) {
	for msg := range client.C {
		conn.Write([]byte(client.Name + msg))
	}
}

func HandleConnect(conn net.Conn) {
	defer conn.Close()
	netAddr := conn.RemoteAddr().String()
	clnt := Client{make(chan string), netAddr, netAddr}
	onlineMap[netAddr] = clnt

	go WriteMessageToClient(clnt, conn)
	//发送上线广播
	message <- "[" + netAddr + "]:" + "online" + "\n"

	buf := make([]byte, 4094)
	for {
		n, _ := conn.Read(buf)
		message <- string(buf[:n])
	}

}

func Manager() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message

		for _, clt := range onlineMap {
			clt.C <- msg
		}
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer listener.Close()

	go Manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err", err)
			return
		}

		go HandleConnect(conn)
	}
}
