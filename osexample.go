package main

import (
    "os"
    "strings"
    "fmt"
)

func main() {
    os.Setenv("WEB" , "http://coolshell.cn")//getenv  , setenv , putenv , unsetenv , 更改环境变量
	fmt.Printf(os.Getenv("WEB"))

    for _ , env := range os.Environ() {
	   e := strings.Split(env , "=")
	   println(e[0] , "=" , e[1])
    }
}
