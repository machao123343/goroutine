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
    "time"
)

const (
    RECV_BUF_LEN = 1024
)

func main() {
	conn , err := net.Dial("tcp" , "127.0.0.1:6666")
    if err != nil {
	   panic(err.Error())
    }
    defer conn.Close()

    buf := make([]byte , RECV_BUF_LEN)

    for i := 0 ; i < 5 ; i++ {
	   //准备要发送到的字符串
	   msg := fmt.Sprintf("Hello World , %03d " , i)
	   n , err := conn.Write([]byte(msg))
	   if err != nil {
		  fmt.Println("Write Buffer Error:" , err.Error())
		  break
	   }
	   fmt.Println(msg)

	   //从服务端接收字符串
	   n , err = conn.Read(buf)
	   if err != nil {
		  fmt.Println("Read Buffer Rrror: " , err.Error())
		  break
	   }
	   fmt.Println(string(buf[0:n]))//接收并输出服务端发出的数据

	   //wait
	   time.Sleep(time.Second)
    }
}
