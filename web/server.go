package main

import (
	"net"
	"fmt"
	"bufio"
)

var ConnMap map[string]*net.TCPConn//定义字典字符型变量

func main() {
	var tcpAddr *net.TCPAddr
	ConnMap = make(map[string]*net.TCPConn)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	//不断接收客户端发来的信息
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		//新连接加入map
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
		//这里返回消息改为广播
		broadcastMessage(conn.RemoteAddr().String() + ":" + string(message))
	}
}

func broadcastMessage(message string) {
	b := []byte(message)
	//遍历所有客户端并发送消息
	for _, conn := range ConnMap {
		conn.Write(b)
	}
}


