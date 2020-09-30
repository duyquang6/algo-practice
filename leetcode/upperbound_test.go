package leetcode

import "testing"

func Test_upperbound(t *testing.T) {
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
			want: 3,
		},
		{
			name: "TC2",
			args: args{
				arr:    []int{10, 20, 20, 30, 30, 40, 50},
				target: 50,
			},
			want: 7,
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
			if got := upperbound(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("upperbound() = %v, want %v", got, tt.want)
			}
		})
	}
}
