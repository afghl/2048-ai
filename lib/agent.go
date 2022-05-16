package lib

import (
	"math"
)

type Agent struct {
	depth       int // how far do we look for the best action
	heuristicFn heuristic
}

func NewAgent(depth int) *Agent {
	return &Agent{depth: depth, heuristicFn: MonotonicSmoothnessHeuristic()}
}

func (a *Agent) GetAction(gameState GameState) Direction {
	act := NONE
	max := math.Inf(-1)
	for action, nextState := range gameState.GenerateDirectionSuccessorState() {
		v := a.value(nextState, false, 0)
		//fmt.Printf("state: %v, score: %v, max is: %v \n", nextState.gridFmt(), v, max)
		if v > max {
			max = v
			act = action
		}
	}
	return act
}

func (a *Agent) value(state GameState, directionMove bool, depth int) float64 {
	//fmt.Printf("current state: %v, directionMove: %v, depth: %v \n", state.gridFmt(), directionMove, depth)
	if depth == a.depth {
		v := a.heuristicFn.Evaluate(state)
		//fmt.Printf("value, current state: %v, score: %v,  \n", state.gridFmt(), v)
		return v
	}
	var value float64
	// if it is a direction action, then select a max value, else select an expected value
	if directionMove {
		value = a.maxValue(state, depth)
	} else {
		value = a.avgValue(state, depth+1)
	}
	return value
}

func (a *Agent) maxValue(gameState GameState, depth int) float64 {
	max := math.Inf(-1)
	for _, nextState := range gameState.GenerateDirectionSuccessorState() {
		v := a.value(nextState, false, depth)
		if v > max {
			max = v
		}
	}
	if max == math.Inf(-1) {
		max = 0.0
	}
	return max
}

func (a *Agent) avgValue(gameState GameState, depth int) float64 {
	sum := 0.0
	tileStates := gameState.GenerateRandomTileSuccessorState()

	for _, state := range tileStates {
		//fmt.Printf("get a value of a tile state: %v, depth: %v \n", state, depth)
		value := a.value(state, true, depth)
		sum += value * state.Weight
	}
	value := sum / float64(len(tileStates))
	//fmt.Printf("avg..depth is: %v, tileState size: %v, sum:%v, value:%v \n", depth, len(tileStates), sum, value)

	return value
}
