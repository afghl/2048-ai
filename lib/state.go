package lib

type GameState struct {
	Grid [][]int
}

func NewState(grid [][]int) GameState {
	return GameState{Grid: grid}
}
