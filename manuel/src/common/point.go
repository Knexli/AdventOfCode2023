package common

import "math"

type Point struct {
	X, Y int
}

type Grid[T any] struct {
	CoordinateSystem map[Point]T
}

func (g Grid[T]) Set(point Point, t T) {
	g.CoordinateSystem[point] = t
}

func (g Grid[T]) Get(point Point) (T, bool) {
	t, i := g.CoordinateSystem[point]
	return t, i
}

func (p Point) Steps(p2 Point) int {
	xDiff := math.Abs(float64(p.X - p2.X))
	yDiff := math.Abs(float64(p.Y - p2.Y))
	return int(yDiff + xDiff)
}
