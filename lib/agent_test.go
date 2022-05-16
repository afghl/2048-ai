package lib

import "testing"

func mockGameState() GameState {
	return GameState{Size: 4, Grid: [][]int{{0, 0, 16, 0}, {0, 4, 0, 2}, {2, 8, 8, 8}, {4, 16, 64, 2}}}
}

func TestAgent_GetAction(t *testing.T) {
	type fields struct {
		depth     int
		evaluator heuristic
	}
	type args struct {
		gameState GameState
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		isNot  Direction
	}{
		{name: "can get an action", fields: fields{depth: 2, evaluator: MonotonicSmoothnessHeuristic()}, args: args{mockGameState()}, isNot: NONE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Agent{
				depth:     tt.fields.depth,
				evaluator: tt.fields.evaluator,
			}
			if got := a.GetAction(tt.args.gameState); got == tt.isNot {
				t.Errorf("GetAction() = %v, isNot %v", got, tt.isNot)
			}
		})
	}
}
