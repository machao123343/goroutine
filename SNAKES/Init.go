package main

import (
	"math/rand"
	"strconv"

	"github.com/nsf/termbox-go"
)


func Random(min, max int) int {
	return rand.Intn(max - min) + min
}

func PrintInt(x, y, val int) {
	for i, c := range strconv.Itoa(val) {//返回固定格式
		termbox.SetCell(x + i, y, c, TextColor, BackgroundColor)
	}
}

func (c *Coord) Draw(color termbox.Attribute) {
	termbox.SetCell(c.x, c.y, ' ', BackgroundColor, color)
}