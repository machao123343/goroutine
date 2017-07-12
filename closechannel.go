package main

import (
    "time"
    "fmt"
    "math/rand"
)

func main(){
    channel := make(chan string)
    rand.Seed(time.Now().Unix())

//向channel发送随机个数的message
	go func() {
   	 cnt := rand.Intn(10)
	    fmt.Println("message cnt : " , cnt)
	    for i := 0 ; i <cnt ; i++ {
		   channel <- fmt.Sprintf("message-%d" , i)
	    }
	    close(channel)
	}()
    var more bool = true
    var msg string
    for more {
	   select{
	   case msg , more = <- channel://bool型判断channel是否开启
		  if more {
			 fmt.Println(msg)
		  } else {
			 fmt.Println("channel closed.")
		  }
	   }
    }
}
