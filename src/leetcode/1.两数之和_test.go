// Copyright (c) 2019 conan17
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"official",
			args{
				[]int{2, 7, 11, 15},
				9,
			},
			[]int{0, 1},
		},
		{
			"the same",
			args{
				[]int{2, 7, 7, 15},
				14,
			},
			[]int{1, 2},
		},
		{
			"interval",
			args{
				[]int{2, 7, 7, 15},
				17,
			},
			[]int{0, 3},
		},
		{
			"error",
			args{
				[]int{2, 7, 7, 15},
				11,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
