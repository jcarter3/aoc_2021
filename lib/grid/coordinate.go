package grid

import (
	"aoc_2021/lib/helper"
	"fmt"
	"log"
	"math"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func ParseCoordinate(input string) *Coordinate {
	var x, y int
	_, err := fmt.Sscanf(input, "%d, %d", &x, &y)
	if err != nil {
		log.Fatalf("Failed to parse coord %s - %v\n", input, err)
	}
	return &Coordinate{
		X: x,
		Y: y,
	}
}

func (c *Coordinate) Move(x, y int) {
	c.X += x
	c.Y += y
}

func (c *Coordinate) MoveLiteral(direction rune, steps int) {
	switch strings.ToUpper(string(direction)) {
	case "R":
		c.X += steps
	case "L":
		c.X -= steps
	case "U":
		c.Y += steps
	case "D":
		c.Y -= steps
	default:
		panic(fmt.Sprintf("unknown direction: %c", direction))
	}
}

func (c *Coordinate) Copy() Coordinate {
	return Coordinate{c.X, c.Y}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%d, %d", c.X, c.Y)
}

func (c *Coordinate) Equal(c2 Coordinate) bool {
	return c.X == c2.X && c.Y == c2.Y
}

func (c *Coordinate) Distance(c2 Coordinate) int {
	x := c.X - c2.X
	y := c.Y - c2.Y
	x = helper.Abs(x)
	y = helper.Abs(y)
	return x + y
}

func (c *Coordinate) Angle(c2 Coordinate) float64 {
	dx := c.X - c2.X
	dy := c.Y - c2.Y
	return math.Atan2(float64(dy), float64(dx))
	//return float64(dx) / float64(dy)
}

func (c *Coordinate) Neighbors(diagonals bool) []Coordinate {
	coords := []Coordinate{
		{c.X, c.Y + 1},
		{c.X, c.Y - 1},
		{c.X + 1, c.Y},
		{c.X - 1, c.Y},
	}
	if diagonals {
		coords = append(coords,
			Coordinate{c.X + 1, c.Y + 1},
			Coordinate{c.X - 1, c.Y - 1},
			Coordinate{c.X + 1, c.Y - 1},
			Coordinate{c.X - 1, c.Y + 1})
	}
	return coords
}

func (c *Coordinate) PathTo(c2 Coordinate) []Coordinate {
	path := make([]Coordinate, 0)
	xStep := 1
	if c.X == c2.X {
		xStep = 0
	} else if c.X > c2.X {
		xStep = -1
	}
	yStep := 1
	if c.Y == c2.Y {
		yStep = 0
	} else if c.Y > c2.Y {
		yStep = -1
	}
	c1 := c.Copy()
	for !c1.Equal(c2) {
		path = append(path, c1.Copy())
		c1.Move(xStep, yStep)
	}
	path = append(path, c2)
	return path
}
