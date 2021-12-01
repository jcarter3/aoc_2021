package grid

type Grid struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (g *Grid) Valid(c Coordinate) bool {
	return c.X >= g.X && c.Y >= g.Y && c.X < g.X+g.Width && c.Y < g.Y+g.Height
}
