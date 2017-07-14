/*
//socker net.Conn类型接口的八种方法


//1 Read基本用法
Read(b []byte) (n int , err error)
b := make([]byte , 10)
n , err := conn.Read(b)
content := string(b[:n])
//reader := bufio.NewReader(conn)
//line , err := reader.ReadBytes('\n')

//2 Read error 处理
var dataBuffer bytes.Buffer
b := make([]byte , 10)
for {
	n , err := conn.Read(b)
	if err != nil {
		if err == io.EOF {
			fmt.Println("The connection is closed.")
			} else {
				fmt.Printf("Read Error : %s\n", err)
			}
			break
		}
		dataBuffer.Write(b[:n])
	}

//3 Write基本用法
write(b []byte) (n int , err error)
writer := bufio.NewWriter(conn)//注意容量

//4 Close 关闭当前连接

//5 LocalAddr和RemoteAddr方法

conn.LocalAddr().Network()//返回本地协议名称
conn.RemoteAddr().String()//返回远端地址

//6 SetDeadline SetReadDeadline SetWriteDeadline方法
//规定超时时间，超时返回定义的error值

//example：Echo Server
*/

package main

import (
    "net"
    "fmt"
    "io"
)

const (
    REVC_BUF_LEN = 1024//初始化缓冲长度
)

func main() {
    listener ,err := net.Listen("tcp" , "0.0.0.0:6666")
    if err != nil {
	   panic("error listening:" + err.Error())
    }
    fmt.Println("Starting the server")
    for {
	   conn , err := listener.Accept()//接受连接
	   if err != nil {
		  panic("Error accept :" +err.Error())
	   }
	   fmt.Println("Accepted the Connection :" , conn.RemoteAddr())
	   go EchoServer(conn)
    }
}

func EchoServer(conn net.Conn) {
    buf := make([]byte , REVC_BUF_LEN)
    defer conn.Close()

    for {
	   n , err := conn.Read(buf)
	   switch err{
	   case nil    :
		  conn.Write(buf[0:n])
	   case io.EOF :
		  fmt.Printf("Warning: End of daa: %s \n" , err)
		  return
	   default:
		  fmt.Printf("Error: Reading data : %s \n" , err)
		  return
	   }
    }
}






