package main

import (
	"fmt"
	"sort"
)


// Definition for singly-linked list.
type ListNode struct {
	Val int
 	Next *ListNode
}


func deleteDuplicates(head *ListNode) *ListNode {

	if head ==nil {
		return nil
	}

	cur := head
	for ;cur.Next!=nil;{
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return head

}



func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
        return nil
    }
    pre := &ListNode {
        Next: head,
    }
    slow, fast := head, head.Next
    for fast != nil {
        if slow.Val == fast.Val {
            slow.Next, fast = fast.Next, fast.Next
        } else {
            slow, fast = fast, fast.Next
        }
    }
    return pre.Next
}