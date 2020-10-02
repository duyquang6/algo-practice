package leetcode

import (
	"reflect"
	"testing"
)

func TestRecentCounter_Ping(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		t []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "tc1",
			fields: fields{
				arr: []int{},
			},
			args: args{
				t: []int{1, 100, 3001, 3002},
			},
			want: []int{1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RecentCounter{
				arr: tt.fields.arr,
			}
			got := []int{}
			for _, arg := range tt.args.t {
				got = append(got, this.Ping(arg))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecentCounter.Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}
