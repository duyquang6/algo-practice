package leetcode

import (
	"strconv"
	"testing"
)

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 2, -2147483648},
			},
			want: -2147483648,
		},
		{
			args: args{
				nums: []int{1, 2, 2, 5, 3, 5},
			},
			want: 2,
		},
	}
	for i, tt := range tests {
		t.Run("TC"+strconv.Itoa(i+1), func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
