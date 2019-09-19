// Copyright (c) 2019 conan17
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package leetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		newVal   int
		carry    bool
		topLevel *ListNode
		current  *ListNode
		root     *ListNode
	)
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			newVal = l1.Val + l2.Val
			l1, l2 = l1.Next, l2.Next
		} else if l1 != nil {
			newVal = l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			newVal = l2.Val
			l2 = l2.Next
		}

		if carry {
			newVal++
			carry = false
		}
		if newVal > 9 {
			newVal -= 10
			carry = true
		}
		if topLevel == nil {
			current = &ListNode{}
			root = current
		} else {
			current = &ListNode{}
			topLevel.Next = current
		}
		current.Val = newVal
		topLevel = current
	}
	if carry {
		current = &ListNode{1, nil}
		topLevel.Next = current
	}

	return root
}
