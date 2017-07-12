package main

import "fmt"

func main() {
    //创建一个string类型的channal
    channel := make(chan string)

    //创建一个goroutine向channel里发一个字符串
    go func() {channel <- "hello" }()

    msg := <- channel
    fmt.Println(msg)
}
