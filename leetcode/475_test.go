// https://leetcode.com/problems/heaters/
package leetcode

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				houses:  []int{1, 2, 3},
				heaters: []int{2},
			},
			want: 1,
		},
		{
			args: args{
				houses:  []int{1, 2, 3, 4},
				heaters: []int{1, 4},
			},
			want: 1,
		},
		{
			args: args{
				houses:  []int{1, 3, 4},
				heaters: []int{1, 4},
			},
			want: 1,
		},
		{
			args: args{
				houses:  []int{1, 10, 20, 50},
				heaters: []int{20},
			},
			want: 30,
		},
		{
			args: args{
				houses:  []int{1, 10, 20, 50},
				heaters: []int{3, 33},
			},
			want: 17,
		},
		{
			args: args{
				houses:  []int{1, 5},
				heaters: []int{10},
			},
			want: 9,
		},
		{
			args: args{
				houses:  []int{1},
				heaters: []int{1, 2, 3, 4},
			},
			want: 0,
		},
		{
			args: args{
				houses:  []int{5},
				heaters: []int{1},
			},
			want: 4,
		},
		{
			args: args{
				houses:  []int{1, 2, 3, 5, 15},
				heaters: []int{2, 30},
			},
			want: 13,
		},
		{
			args: args{
				houses:  []int{1, 1, 1, 1, 1, 1, 999, 999, 999, 999, 999},
				heaters: []int{499, 500, 501},
			},
			want: 498,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
