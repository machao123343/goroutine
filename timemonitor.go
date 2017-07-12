package main

import "time"

func main() {
    for i := 0 ; i < 10 ; i++ {
	   //创建10个goroutine
	   go func() {
		  time.Sleep(5 * time.Second)
		  //每个睡眠5秒
	   }()
	   time.Sleep(time.Second)//主程睡眠1秒
    }
}
