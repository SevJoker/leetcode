package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func Print(list1 *ListNode) {

	for list1 != nil {
		fmt.Println("ListNode.Val:",list1.Val)
		list1 = list1.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmp := &ListNode{}
	head := tmp
	for list2 != nil && list1 != nil {
		if list1.Val > list2.Val {
			head.Next = list2
			list2 = list2.Next
		}else{
			head.Next = list1
			list1 = list1.Next
		}
		head = head.Next
	}

	for list1 != nil {
		head.Next = list1
		list1 = list1.Next
		head = head.Next
	}

	for list2 != nil {
		head.Next = list2
		list2 = list2.Next
		head = head.Next
	}
	return tmp.Next
}


func main() {
	l1 := &ListNode{Val:1}
	l1.Next = &ListNode{Val:2}
	l1.Next.Next = &ListNode{Val:6}
	l1.Next.Next.Next = &ListNode{Val:7}
	l1.Next.Next.Next.Next = &ListNode{Val:9}

	l2 := &ListNode{Val:3}
	l2.Next = &ListNode{Val:4}
	l2.Next.Next = &ListNode{Val:10}


	Print(mergeTwoLists(l1,l2))
}