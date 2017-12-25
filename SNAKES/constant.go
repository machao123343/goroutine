package main

import (
	"time"
	"github.com/nsf/termbox-go"
)

const (
	Speed = 50 * time.Millisecond
	GrowAmount = 10
	FoodCount  = 5
	TextColor  = termbox.ColorBlue
	BackgroundColor = termbox.ColorDefault
	SnakeColor  = termbox.ColorBlue
	FoodColor = termbox.ColorBlue
)

type Direction int

const (
	Up Direction = iota//对下采取相同操作
	Down
	Left
	Right
)

type Coord struct {
	x, y int
}

type Snake struct {
	direction Direction
	body      []*Coord
	coords    map[Coord]bool
	grow      int

}

type Context struct {
	quit     bool
	snake    *Snake
	foods    map[Coord]bool//二维字典
	skip     bool
}

