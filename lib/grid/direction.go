package grid

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

// North & South reversed since grids are commonly based from upper left
var directions = []Coordinate{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type Direction struct {
	curDir int
}

func (d *Direction) Deltas() (int, int) {
	return directions[d.curDir].X, directions[d.curDir].Y
}

func (d *Direction) Left() Direction {
	dir := (d.curDir + 3) % 4
	return Direction{dir}
}

func (d *Direction) Right() Direction {
	dir := (d.curDir + 1) % 4
	return Direction{dir}
}

func (d *Direction) String() string {
	switch d.curDir {
	case NORTH:
		return "NORTH"
	case EAST:
		return "EAST"
	case SOUTH:
		return "SOUTH"
	case WEST:
		return "WEST"
	}
	return "uhh"
}

func NewDirection(curDir int) Direction {
	return Direction{curDir}
}
