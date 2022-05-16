package lib

import (
	"reflect"
	"testing"
)

func Test_move(t *testing.T) {
	type args struct {
		grid      [][]int
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "up", args: args{
			grid: [][]int{
				{2, 4, 0, 0}, {4, 2, 0, 2}, {16, 2, 16, 4}, {2, 64, 8, 2},
			}, direction: UP,
		}, want: [][]int{
			{2, 4, 16, 2}, {4, 4, 8, 4}, {16, 64, 0, 2}, {2, 0, 0, 0},
		},
		},
		{name: "down", args: args{
			grid: [][]int{
				{2, 4, 16, 2}, {2, 8, 8, 4}, {0, 16, 64, 4}, {0, 0, 0, 2},
			}, direction: DOWN,
		}, want: [][]int{
			{0, 0, 0, 0}, {0, 4, 16, 2}, {0, 8, 8, 8}, {4, 16, 64, 2},
		},
		},
		{name: "right", args: args{
			grid: [][]int{
				{2, 4, 16, 2}, {4, 4, 8, 4}, {16, 64, 2, 2}, {2, 0, 0, 0},
			}, direction: RIGHT,
		}, want: [][]int{
			{2, 4, 16, 2}, {0, 8, 8, 4}, {0, 16, 64, 4}, {0, 0, 0, 2},
		},
		},
		{name: "left", args: args{
			grid: [][]int{
				{0, 0, 0, 0}, {0, 4, 16, 2}, {2, 8, 8, 8}, {4, 16, 64, 2},
			}, direction: LEFT,
		}, want: [][]int{
			{0, 0, 0, 0}, {4, 16, 2, 0}, {2, 16, 8, 0}, {4, 16, 64, 2},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.grid, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, isNot %v", got, tt.want)
			}
		})
	}
}
