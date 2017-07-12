package main

import (
    "time"
    "fmt"
    "math/rand"
)

func routine(name string, delay time.Duration) {

    t0 := time.Now()
    fmt.Println(name, " start at ", t0)

    time.Sleep(delay)

    t1 := time.Now()
    fmt.Println(name, " end at ", t1)

    fmt.Println(name, " lasted ", t1.Sub(t0))
}

func main() {

    //生成随机种子
    rand.Seed(time.Now().Unix())

    var name string
    for i:=0; i<3; i++{
	   name = fmt.Sprintf("go_%02d", i) //生成ID
	   //生成随机等待时间，从0-4秒
	   go routine(name, time.Duration(rand.Intn(5)) * time.Second)
    }

    //让主进程停住，不然主进程退了，goroutine也就退了
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
