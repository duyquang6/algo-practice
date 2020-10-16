package leetcode

import (
	"reflect"
	"testing"
)

func Test_mostVisited(t *testing.T) {
	type args struct {
		n      int
		rounds []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:      4,
				rounds: []int{1, 3, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			args: args{
				n:      2,
				rounds: []int{2, 1, 2, 1, 2, 1, 2, 1, 2},
			},
			want: []int{2},
		},
		{
			args: args{
				n:      7,
				rounds: []int{1, 3, 5, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostVisited(tt.args.n, tt.args.rounds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostVisited() = %v, want %v", got, tt.want)
			}
			if got := optimizedMostVisited(tt.args.n, tt.args.rounds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optimizedMostVisited() = %v, want %v", got, tt.want)
			}
		})
	}
}
