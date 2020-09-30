package leetcode

import "testing"

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
				capacity: 4,
			},
			want: false,
		},
		{
			name: "tc2",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
				capacity: 5,
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
				capacity: 3,
			},
			want: true,
		},
		{
			name: "tc4",
			args: args{
				trips:    [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
				capacity: 11,
			},
			want: true,
		},
		{
			name: "tc5",
			args: args{
				trips:    [][]int{{3, 1, 5}, {3, 3, 7}, {8, 4, 6}},
				capacity: 13,
			},
			want: false,
		},
		{
			name: "tc6",
			args: args{
				trips:    [][]int{{3, 1, 5}, {3, 3, 7}, {8, 4, 6}},
				capacity: 14,
			},
			want: true,
		},
		{
			name: "tc7",
			args: args{
				trips:    [][]int{{9, 3, 4}, {9, 1, 7}, {4, 2, 4}, {7, 4, 5}},
				capacity: 23,
			},
			want: true,
		},
		{
			name: "tc8",
			args: args{
				trips:    [][]int{{10, 4, 5}, {8, 1, 2}, {3, 1, 8}, {4, 2, 8}, {3, 4, 5}, {1, 1, 9}, {1, 1, 5}, {2, 4, 7}},
				capacity: 23,
			},
			want: false,
		},
		{
			name: "tc9",
			args: args{
				trips:    [][]int{{6, 6, 7}, {5, 1, 5}, {1, 4, 7}, {6, 3, 7}, {1, 1, 3}, {3, 1, 7}, {8, 4, 6}, {10, 2, 9}, {8, 6, 7}},
				capacity: 33,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
