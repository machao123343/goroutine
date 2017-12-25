package main

import (
	"time"
	"fmt"
	"math/rand"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	ctx := NewContext()
	rand.Seed(time.Now().Unix())

	for !ctx.quit {
		ctx.Update()
		select {
		case e := <-events:
			switch e.Type {
			case termbox.EventKey:
				if e.Key == termbox.KeyEsc {
					ctx.quit = true
				} else {
					ctx.Handlekey(e.Key)
				}
			}
		default:
			ctx.Draw()
			time.Sleep(Speed)
		}
	}

	termbox.Close()
	fmt.Println("Game over ! Your score is ", ctx.Score())
}