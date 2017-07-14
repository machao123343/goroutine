package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    names := []string{"Eric" , "Harry" , "Robert" , "Jim" , "Mark"}
    wg := sync.WaitGroup{}
    for _ , name := range names {
	   wg.Add(1)
	   go func(who string) {
		  runtime.Gosched()
		  fmt.Printf("Hello , %s \n" , who)
	   wg.Done()
	   }(name)
    }
    wg.Wait()
    //time.Sleep(n * time.Second)延长kill子进程的时间，达到输出的目的
}

/*
package main

import (
    "runtime"
)

func main() {
    go println("Go! Goroutine!")
    runtime.Gosched()
}

package main

import (
    "fmt"
    "runtime"
)

func main() {
    name := "Eric"
    go func() {
	   fmt.Printf("Hello , %s.\n", name)
    }()
    name = "Harry"
    runtime.Gosched()
}
//先Harry再运行go func



package main

import (
    "fmt"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(5)
    names := []string{"Eric" , "Harry" , "Robert" , "Jim" , "Mark"}
    for _ , name := range names {
	   go func() {
		  fmt.Printf("Hello , %s \n" , name)
	   }()
	   runtime.Gosched()
    }

}
//不稳定，不能稳定输出固定数和固定顺序
*/