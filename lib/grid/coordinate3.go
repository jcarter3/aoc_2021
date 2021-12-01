package grid

import (
	"fmt"

	"math"

	"gitlab.com/jscarter3/aoc_go/common/helper"
)

type Coordinate3 struct {
	X int
	Y int
	Z int
}

func (c *Coordinate3) Move(x, y, z int) {
	c.X += x
	c.Y += y
	c.Z += z
}

func (c *Coordinate3) Copy() Coordinate3 {
	return Coordinate3{c.X, c.Y, c.Z}
}

func (c *Coordinate3) Values() (int, int, int) {
	return c.X, c.Y, c.Z
}

func (c *Coordinate3) String() string {
	return fmt.Sprintf("X: %d, Y: %d, Z: %d", c.X, c.Y, c.Z)
}

func (c *Coordinate3) Distance(c2 Coordinate3) int {
	x := c.X - c2.X
	y := c.Y - c2.Y
	z := c.Z - c2.Z
	x = helper.Abs(x)
	y = helper.Abs(y)
	z = helper.Abs(z)
	return x + y + z
}

func (c *Coordinate3) Magnitude() float64 {
	return math.Sqrt(float64(c.X*c.X + c.Y*c.Y + c.Z*c.Z))
}

func (c *Coordinate3) Add(c2 Coordinate3) {
	c.X += c2.X
	c.Y += c2.Y
	c.Z += c2.Z
}
