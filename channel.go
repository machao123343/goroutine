package main

import (
    "time"
    "fmt"
   // "sync"
)

func main() {
    ch := make(chan int , 5)
    //wg1 := sync.WaitGroup{}
    //wg2 := sync.WaitGroup{}
    sign := make(chan byte , 2)
    //wg1.Add(1)
    //wg2.Add(1)
    go func() {

	   for i := 0 ; i < 5 ; i++ {

		  ch <- i
		  time.Sleep(2 * time.Second)

	   }
	   //wg1.Done()
	   close(ch)
	   fmt.Println("The channel is colsed.")
	   sign <- 0
    }()
    //wg1.Wait()


    go func() {

	   for{
		  e , ok := <- ch
		  fmt.Printf("%d(%v)\n" , e , ok)
		  if !ok {
			 break
		  }
		  time.Sleep(2*time.Second)
	   }
	   fmt.Println("Done !")
	//   wg2.Done()
    }()
    //wg2.Wait()
	<- sign//延时分Goroutine
     <- sign//延时主Goroutine

}
