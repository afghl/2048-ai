package utils

import (
	"reflect"
	"testing"
)

func Test_deepCopy(t *testing.T) {
	type args struct {
		ori [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "success", args: args{ori: [][]int{{1, 2, 3}, {2, 3, 4}, {7, 8}, {11, 12, 13, 14}}}, want: [][]int{{1, 2, 3}, {2, 3, 4}, {7, 8}, {11, 12, 13, 14}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepCopy(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
