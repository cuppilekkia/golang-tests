package utils

import (
	"reflect"
	"testing"
)

func TestBSort(t *testing.T) {
	type args struct {
		els []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"testing list", args{[]int{6, 2, 4, 7, 1}}, []int{1, 2, 4, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BSort(tt.args.els); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBSort10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	}
}
