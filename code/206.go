package main

import (
	"fmt"
)


type ListNode struct {
    Val int
 	Next *ListNode
 }

func reverse(head *ListNode) *ListNode {
	var (
		cur = head
		pre *ListNode = nil
	)


	for ;cur != nil; {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
    return pre
}

func main() {
	h := &ListNode{Val:1}
	h.Next = &ListNode{Val:2}
	h.Next.Next = &ListNode{Val:3}
	h.Next.Next.Next = &ListNode{Val:4}

	r := reverse(h)
	for ;r!=nil; {
		fmt.Println("-------",r)
		r = r.Next	
	}

}