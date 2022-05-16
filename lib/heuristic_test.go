package lib

import "testing"

func monotonicGameState() GameState {
	return GameState{Size: 4, Grid: [][]int{
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{16, 4, 0, 0},
		{8, 16, 2, 0}}}

	//	_ _ _ _
	//2 _ _ _
	//16 4 _ _
	//8 16 2 _
}

func monotonicState2() GameState {
	//	_ _ _ 2
	//_ _ 16 4
	//_ _ _ 4
	//_ 4 16 2
	return GameState{Size: 4, Grid: [][]int{
		{0, 0, 0, 2},
		{0, 0, 16, 4},
		{0, 0, 0, 4},
		{0, 4, 16, 2}}}
}

func monotonicState3() GameState {
	//_ _ _ _
	//2 _ 4 _
	//32 2 256 _
	//8 128 2 2
	return GameState{Size: 4, Grid: [][]int{
		{0, 0, 0, 0},
		{2, 0, 4, 0},
		{32, 2, 256, 0},
		{8, 128, 2, 2}}}
}

func Test_monotonicSmoothnessEvaluator_getMonotonic(t *testing.T) {
	type args struct {
		state GameState
	}
	tests := []struct {
		name  string
		args  args
		up    float64
		down  float64
		left  float64
		right float64
	}{
		{name: "get result", args: args{monotonicGameState()}, up: 1.0, down: 9.0, left: 9.0, right: 1.0},
		{name: "get result success", args: args{monotonicState2()}, up: 1.0, down: 7.0, left: 5.0, right: 11.0},
		{name: "get result success", args: args{monotonicState3()}, up: 9.0, down: 21.0, left: 20.0, right: 12.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := getMonotonic(tt.args.state)
			if got != tt.up {
				t.Errorf("getMonotonic() got = %v, up %v", got, tt.up)
			}
			if got1 != tt.down {
				t.Errorf("getMonotonic() got1 = %v, up %v", got1, tt.down)
			}
			if got2 != tt.left {
				t.Errorf("getMonotonic() got2 = %v, up %v", got2, tt.left)
			}
			if got3 != tt.right {
				t.Errorf("getMonotonic() got3 = %v, up %v", got3, tt.right)
			}
		})
	}
}
