package main

//lint:ignore ST1001 dot import
import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMiddle(head)
	r := mid.Next
	mid.Next = nil

	ll := sortList(head)
	lr := sortList(r)

	return merge(ll, lr)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

// @lc code=end
