package datastruct

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkedList(input string) *ListNode {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil
	}

	var parts = strings.Split(input, ",")

	var head, cur *ListNode
	for _, v := range parts {
		vi, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Atoi error")
			panic(err)
		}
		node := &ListNode{vi, nil}
		if head == nil {
			head = node
		} else {
			cur.Next = node
		}
		cur = node
	}

	return head
}

func PrintLinkedList(head *ListNode) {
	if head == nil {
		fmt.Println("")
	}

	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(", ")
		}
		head = head.Next
	}
	fmt.Println()
}
