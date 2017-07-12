package main

import (
	"fmt"
	"time"
)

func main() {
	//创建两个channel
	c1 := make(chan string)
	c2 := make(chan string)

	//创建两个goroutine来分别向两个channel发送数据
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "hello"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "world"
	}()

	//使用select来侦听两个channel
	/*   for i := 0 ; i < 2 ; i++ {
		   select {
		   case msg1 := <-c1:
			  fmt.Println("received" , msg1)
		   case msg2 := <-c2:
			  fmt.Println("received" , msg2)
		   }
	    }
	}
	*/
    //避免select阻塞
	for {
		timeout_cnt := 0
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 30):
			fmt.Println("time out")
			timeout_cnt++
		}
		if timeout_cnt > 3 {
			break
		}
	}
    /*for {
    		select {
    		case msg1 := <-c1:
    			fmt.Println("received" , msg1)
    		case msg2 := <-c2:
    			fmt.Println("received" , msg2)
    		default :
    			fmt.Println("received" , msg2)
    			time.Sleep(time.Second)
    		}
    	}
     */
    //另一个无阻塞方式
}
