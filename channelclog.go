package main

import (
    "fmt"
    "time"
)

func main() {
    channel := make(chan string) //注意： buffer为1

    go func() {
	   channel <- "hello"
	   fmt.Println("write \" hello\" done!")

	   channel <- "world" //Reader在Sleep，这里在堵塞
	   fmt.Println("write \"world\" done!")

	   fmt.Println("write go sleep ...")
	   time.Sleep(3 * time.Second)
	   channel <- "channel"
	   fmt.Println("Write \" channel \" done!")
    }()

    time.Sleep(2 * time.Second)
    fmt.Println("Reader Wake up...")

    msg := <- channel
    fmt.Println("Reader: " , msg)

    msg = <- channel
    fmt.Println("Reader: " , msg)

    msg = <- channel
    fmt.Println("Reader: " , msg)
}
