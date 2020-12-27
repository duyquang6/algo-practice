package leetcode

import "testing"

func Test_maxRepeating(t *testing.T) {
	type args struct {
		sequence string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				sequence: "ababc",
				word:     "ab",
			},
			want: 2,
		},
		{
			args: args{
				sequence: "ababc",
				word:     "ba",
			},
			want: 1,
		},
		{
			args: args{
				sequence: "ababc",
				word:     "ac",
			},
			want: 0,
		},
		{
			args: args{
				sequence: "aaabaaaabaaabaaaabaaaabaaaabaaaaba",
				word:     "aaaba",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeating(tt.args.sequence, tt.args.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
