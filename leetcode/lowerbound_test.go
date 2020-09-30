package leetcode

import "testing"

func Test_lowerbound(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TC1",
			args: args{
				arr:    []int{10, 20, 20, 30, 30, 40, 50},
				target: 20,
			},
			want: 1,
		},
		{
			name: "TC2",
			args: args{
				arr:    []int{10, 20, 20, 30, 30, 40, 50},
				target: 50,
			},
			want: 6,
		},
		{
			name: "TC3",
			args: args{
				arr:    []int{10, 20, 20, 30, 30, 40, 50},
				target: 60,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowerbound(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("lowerbound() = %v, want %v", got, tt.want)
			}
		})
	}
}
