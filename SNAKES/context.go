package main

import (
	"github.com/nsf/termbox-go"
)

func NewContext() *Context {
	w, h := termbox.Size()
	return &Context {
		snake: NewSnake(w/2, h/2),
		foods: make(map[Coord]bool),
	}
}

func (ctx *Context) Grow(s *Snake) {
	w, h := termbox.Size()
	head := s.Head()
	c := &Coord{head.x, head.y}
	switch s.direction {
	case Up:
		c.y--
		if c.y < 0 {
			c.y = h - 1
		}
	case Down:
		c.y++
		if c.y >= h {
			c.y = 0
		}
	case Left:
		c.x--
		if c.x < 0 {
			c.x = w - 1
		}
	case Right:
		c.x++
		if c.x >= w {
			c.x = 0
		}
	}
	if s.Occupies(c) {
		ctx.quit = true
	} else {
		s.Push(c)
	}
}

func (ctx *Context) Move(s *Snake) {
	ctx.Grow(s)
	if s.grow <= 0 {
		s.Pop()
	} else {
		s.grow--
	}
}

func (ctx *Context) Score() int {
	return len(ctx.snake.body) - 1
}

func (ctx *Context) DrawFoods() {
	for food := range ctx.foods {//随地遍历食物点
		food.Draw(FoodColor)
	}
}

func (ctx *Context) AddFoods() {
	w, h := termbox.Size()
	for len(ctx.foods) < FoodCount {
		food := Coord{Random(0, w - 1), Random(0, h - 1)}
		ctx.foods[food] = true
	}
}

func (ctx *Context) Draw() {
	termbox.Clear(BackgroundColor, BackgroundColor)
	PrintInt(0, 0, ctx.Score())
	ctx.DrawFoods()
	ctx.snake.Draw()
	termbox.Flush()//刷新
}

func (ctx *Context) Update() {
	if ctx.snake.direction == Up || ctx.snake.direction == Down {
		ctx.skip = !ctx.skip
		if ctx.skip {
			return
		}
	} else {
		ctx.skip = false
	}
	ctx.Move(ctx.snake)
	for food := range ctx.foods {
		if ctx.snake.Occupies(&food) {
			ctx.snake.grow += GrowAmount
			delete(ctx.foods, food)
		}
	}
	ctx.AddFoods()
}

func (ctx *Context) Handlekey(key termbox.Key) {
	switch key {
	case termbox.KeyArrowUp:
		if ctx.snake.direction != Down {
			ctx.snake.direction = Up
		}
	case termbox.KeyArrowDown:
		if ctx.snake.direction != Up {
			ctx.snake.direction = Down
		}
	case termbox.KeyArrowLeft:
		if ctx.snake.direction != Right {
			ctx.snake.direction = Left
		}
	case termbox.KeyArrowRight:
		if ctx.snake.direction != Left {
			ctx.snake.direction = Right
		}
	}
}
