// Copyright (c) 2019 conan17
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
func twoSum(nums []int, target int) []int {
	cache := map[int]int{}
	for i := range nums {
		sub := target - nums[i]
		if v, ok := cache[sub]; ok && i != v {
			return []int{v, i}
		}
		cache[nums[i]] = i
	}
	return []int{}
}
