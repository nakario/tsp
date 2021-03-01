package tsp

// Vector describes a 2D vector.
type Vector struct {
        X, Y int
}

// Origin is the origin of the xy plane.
var Origin = Point{X: 0, Y: 0}

// Point describes a 2D point.
type Point struct {
        X, Y int
}

// Add returns a new 2D point p + v
func (p Point) Add(v Vector) Point {
        return Point{X: p.X + v.X, Y: p.Y + v.Y}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ChebyshevDistance calculates the Chebyshev distance between 2 points.
func (p Point) ChebyshevDistance(p2 Point) int {
	return max(abs(p2.X - p.X), abs(p2.Y - p.Y))
}

// ManhattanDistance calculates the Manhattan distance between 2 points.
func (p Point) ManhattanDistance(p2 Point) int {
	return abs(p2.X - p.X) + abs(p2.Y - p.Y)
}

// Area describes parts of the xy plane.
type Area int

const (
	// AreaOnLine contains points on the x-axis, y-axis, the line y = x and the line y = -x.
	AreaOnLine Area = iota
	AreaENE
	AreaNNE
	AreaNNW
	AreaWNW
	AreaWSW
	AreaSSW
	AreaSSE
	AreaESE
)

// Area returns which area the point p belongs to.
func (p Point) Area() Area {
	if p.X > 0 && p.Y > 0 {
		if p.X > p.Y {
			return AreaENE
		} else if p.X < p.Y {
			return AreaNNE
		}
	} else if p.X < 0 && p.Y > 0 {
		if -p.X < p.Y {
			return AreaNNW
		} else if -p.X > p.Y {
			return AreaWNW
		}
	} else if p.X < 0 && p.Y < 0 {
		if -p.X > -p.Y {
			return AreaWSW
		} else if -p.X < -p.Y {
			return AreaSSW
		}
	} else if p.X > 0 && p.Y < 0 {
		if p.X < -p.Y {
			return AreaSSE
		} else if p.X > -p.Y {
			return AreaESE
		}
	}
	return AreaOnLine
}
