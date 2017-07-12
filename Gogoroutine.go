package main

import "fmt"

func f(msg string) {
    fmt.Printf(msg)
}

func main() {
    go f("goroutine")
    go func(msg string) {
	   fmt.Printf(msg)
    }("going")
}
