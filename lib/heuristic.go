package lib

import (
	"math"
)

type heuristic interface {
	Evaluate(state GameState) float64
}

func MonotonicSmoothnessHeuristic() heuristic {
	return &monotonicSmoothnessEvaluator{}
}

type monotonicSmoothnessEvaluator struct {
}

func (b *monotonicSmoothnessEvaluator) Evaluate(state GameState) float64 {
	um, dm, lm, rm := getMonotonic(state)
	monotonic := math.Min(um, dm) + math.Min(lm, rm)
	emptyCells := emptyCellsCount(state)
	return math.Log(emptyCells)*2.7 - monotonic*0.1
}

func getMonotonic(state GameState) (float64, float64, float64, float64) {
	grid := state.Grid
	size := state.Size
	um, dm, lm, rm := 0.0, 0.0, 0.0, 0.0

	// get vertical monotonic
	for x := 0; x < size; x++ {
		current, next := 0, 1
		for next < size {
			// ignore zero
			for next < size && grid[x][next] == 0 {
				next++
			}
			if next >= size {
				next--
			}
			currentVal, nextVal := grid[x][current], grid[x][next]
			// left > right
			if currentVal > nextVal {
				lm += logValue(currentVal) - logValue(nextVal)
			} else if currentVal < nextVal {
				rm += logValue(nextVal) - logValue(currentVal)
			}
			current = next
			next++
		}
	}
	// get horizontal monotonic
	for y := 0; y < size; y++ {
		current, next := 0, 1
		for next < size {
			// ignore zero
			for next < size && grid[next][y] == 0 {
				next++
			}
			if next >= size {
				next--
			}
			currentVal, nextVal := grid[current][y], grid[next][y]
			// up > down
			if currentVal > nextVal {
				um += logValue(currentVal) - logValue(nextVal)
			} else if currentVal < nextVal {
				dm += logValue(nextVal) - logValue(currentVal)
			}
			current = next
			next++
		}
	}
	return um, dm, lm, rm
}

func logValue(v int) float64 {
	if v == 0 {
		return 0
	} else {
		return math.Log(float64(v)) / math.Log(2.0)
	}
}

type emptyHeuristic struct {
}

func EmptyCellCountHeuristic() heuristic {
	return &emptyHeuristic{}
}

func (b *emptyHeuristic) Evaluate(s GameState) float64 {
	return emptyCellsCount(s)
}

func emptyCellsCount(s GameState) float64 {
	emptyCells := 0.0
	for i := 0; i < len(s.Grid); i++ {
		for j := 0; j < len(s.Grid[i]); j++ {
			if s.Grid[i][j] == 0 {
				// can add new tile in this grid
				emptyCells += 1
			}
		}
	}
	return emptyCells
}
