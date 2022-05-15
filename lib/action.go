package lib

type Action int

var (
	UP    Action = 0
	RIGHT Action = 1
	DOWN  Action = 2
	LEFT  Action = 3
)

var actions = []Action{UP, RIGHT, DOWN, LEFT}
