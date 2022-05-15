package lib

import "testing"

func mockGameState() GameState {
	return GameState{Size: 4, Grid: [][]int{{0, 0, 0, 0}, {0, 4, 16, 2}, {2, 8, 8, 8}, {4, 16, 64, 2}}}
}

func TestAgent_GetAction(t *testing.T) {
	type fields struct {
		depth     int
		evaluator Evaluator
	}
	type args struct {
		gameState GameState
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Direction
	}{
		{name: "can get an action", fields: fields{depth: 2, evaluator: BaseEvaluator()}, args: args{mockGameState()}, want: LEFT},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Agent{
				depth:     tt.fields.depth,
				evaluator: tt.fields.evaluator,
			}
			if got := a.GetAction(tt.args.gameState); got != tt.want {
				t.Errorf("GetAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
