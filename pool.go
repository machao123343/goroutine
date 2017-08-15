package main

import ("runtime/debug"
		"sync/atomic"
		"sync"

	"fmt"
	"runtime"
)

func main() {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count  int32

	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	//返回一个接口型的变量
	pool := sync.Pool{New: newFunc}
	//临时池必须存入一个接口型变量

	v1 := pool.Get()
	fmt.Println("v1: %v\n", v1)

	pool.Put(newFunc())
	pool.Put(newFunc())
	pool.Put(newFunc())
    v2 := pool.Get()
	fmt.Printf("v2: %v\n", v2)

	debug.SetGCPercent(100)
	runtime.GC()
	//把各本地池中的物品回收起来。
	v3 := pool.Get()
	fmt.Printf("v3: %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("v4: %v\n", v4)
}