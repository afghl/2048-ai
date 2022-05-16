package lib

import (
	"math"
)

type Agent struct {
	depth     int // how far do we look for the best action
	evaluator Evaluator
}

func NewAgent(depth int) *Agent {
	return &Agent{depth: depth, evaluator: BaseEvaluator()}
}

func (a *Agent) GetAction(gameState GameState) Direction {
	act := NONE
	max := math.MinInt
	for action, nextState := range gameState.GenerateDirectionSuccessorState() {
		//fmt.Printf("GetAction..next state is: %v, action: %v\n", nextState.Grid, action)
		if v := a.value(nextState, false, 0); v > max {
			max = v
			act = action
		}
	}
	return act
}

func (a *Agent) value(state GameState, directionMove bool, depth int) int {
	//fmt.Printf("current state: %v, directionMove: %v, depth: %v \n", state.gridFmt(), directionMove, depth)
	if depth == a.depth {
		return a.evaluator.Evaluate(state)
	}
	var value int
	// if it is a direction action, then select a max value, else select an expected value
	if directionMove {
		value = a.maxValue(state, depth)
	} else {
		value = a.avgValue(state, depth+1)
	}
	return value
}

func (a *Agent) maxValue(gameState GameState, depth int) int {
	max := math.MinInt
	for _, nextState := range gameState.GenerateDirectionSuccessorState() {
		//fmt.Printf("maxValue..next state is: %v, action: %v\n", nextState.Grid, Direction(action))
		if v := a.value(nextState, false, depth); v > max {
			max = v
		}
	}
	return max
}

func (a *Agent) avgValue(gameState GameState, depth int) int {
	sum := float32(0)
	tileStates := gameState.GenerateRandomTileSuccessorState()

	for _, state := range tileStates {
		//fmt.Printf("get a value of a tile state: %v, depth: %v \n", state, depth)
		value := a.value(state, true, depth)
		sum += float32(value) * state.Weight
	}
	value := int(sum / float32(len(tileStates)))
	//fmt.Printf("avg..depth is: %v, tileState size: %v, sum:%v, value:%v \n", depth, len(tileStates), sum, value)
	return value
}
