package leetcode

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				A: "ab",
				B: "ba",
			},
			want: true,
		},
		{
			args: args{
				A: "ab",
				B: "ab",
			},
			want: false,
		},
		{
			args: args{
				A: "aa",
				B: "aa",
			},
			want: true,
		},
		{
			args: args{
				A: "aaaaaaabc",
				B: "aaaaaaacb",
			},
			want: true,
		},
		{
			args: args{
				A: "",
				B: "aa",
			},
			want: false,
		},
		{
			args: args{
				A: "abcaa",
				B: "abcbb",
			},
			want: false,
		},
		{
			args: args{
				A: "abab",
				B: "abab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("buddyStrings(%v, %v) = %v, want %v", tt.args.A, tt.args.B, got, tt.want)
			}
		})
	}
}
