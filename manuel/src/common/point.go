package common

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
