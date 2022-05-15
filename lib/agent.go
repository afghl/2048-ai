package lib

import "math"

type Agent struct {
	depth     int // how far do we look for the best action
	evaluator Evaluator
}

func NewAgent() *Agent {
	return &Agent{depth: 3, evaluator: BaseEvaluator()}
}

func (a *Agent) GetAction(gameState GameState) Direction {
	act := NONE
	max := math.MinInt
	for _, action := range gameState.GetLegalActions() {
		nextState := gameState.SuccessorState(action)
		v := a.value(nextState, false, 0)
		if v > max {
			max = v
			act = action
		}
	}
	return act
}

func (a *Agent) value(state GameState, arrowMove bool, depth int) int {
	if depth == a.depth {
		return a.evaluator.Evaluate(state)
	}
	var value int
	if arrowMove {
		value = a.maxValue(state, depth)
	} else {
		value = a.avgValue(state, depth-1)
	}
	return value
}

func (a *Agent) maxValue(gameState GameState, depth int) int {
	max := math.MinInt
	for _, action := range gameState.GetLegalActions() {
		nextState := gameState.SuccessorState(action)
		if v := a.value(nextState, true, depth); v > max {
			max = v
		}
	}
	return max
}

func (a *Agent) avgValue(gameState GameState, depth int) int {
	sum := float32(0)
	tileStates := gameState.GenerateRandomTileState()

	for _, state := range tileStates {
		value := a.value(state, false, depth)
		sum += float32(value) * state.Weight
	}
	return int(sum / float32(len(tileStates)))
}
