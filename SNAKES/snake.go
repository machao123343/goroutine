package main

func NewSnake(x, y int) *Snake {
	snake := &Snake{
		direction: Up,
		body:      make([]*Coord, 0),
		coords:    make(map[Coord]bool),
		grow:      GrowAmount,
	}
	snake.Push(&Coord{x, y})
	return snake
}

func (s *Snake) Draw() {
	for _, c := range s.body {
		c.Draw(SnakeColor)
	}
}

func (s *Snake) Push(c *Coord) {//粘贴增加长度
	s.body = append(s.body, c)
	s.coords[*c] = true
}

func (s *Snake) Pop() {
	delete(s.coords, *s.body[0])
	s.body = s.body[1:]
}

func (s *Snake) Head() *Coord{
	return s.body[len(s.body) - 1]
}

func (s *Snake) Occupies(c *Coord) bool {
	return s.coords[*c]
}
