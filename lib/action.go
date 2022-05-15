package lib

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT

	NONE Direction = -1
)

var actions = []Direction{UP, RIGHT, DOWN, LEFT}
