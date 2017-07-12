package main

import (
    "time"
    "fmt"
    "math/rand"
    "runtime"
    "sync"
)

var total_tickets int32 = 10
var mutex = &sync.Mutex{} //可简写成：var mutex sync.Mutex

func sell_tickets(i int){
    for total_tickets>0 {
	   mutex.Lock()
	   if total_tickets > 0 {
		  time.Sleep( time.Duration(rand.Intn(5)) * time.Millisecond)
		  total_tickets--
		  fmt.Println("id : " , i , "ticket : " , total_tickets)
	   }
	   mutex.Unlock()
    }
}
/*
func sell_tickets(i int){
    for{
        if total_tickets > 0 { //如果有票就卖
            time.Sleep( time.Duration(rand.Intn(5)) * time.Millisecond)
            total_tickets-- //卖一张票
            fmt.Println("id:", i, "  ticket:", total_tickets)
        }else{
            break
        }
    }
}
*/
//上面的函数有并发安全性危险

func main() {
    runtime.GOMAXPROCS(4)
    rand.Seed(time.Now().Unix())//生成随机种子

    for i := 0 ; i < 5 ; i++ {//并发5个goroutine来买票
	   go sell_tickets(i)
    }
    //等待线程执行完
    var input string
    fmt.Scanln(&input)
    fmt.Println(total_tickets , "done") //退出打印时还有多少票
}


