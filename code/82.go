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
	headTmp := ListNode{Next:head}
	n1 := headTmp
	n2 := head


	for ;n2!=nil; {
		if n2.Next !=nil && n2.Val != n2.Next.Val {
			n1 = n1.Next
			n2 = n2.Next
		}else{
			n2Tmp := n2.Next
			for ;n2Tmp.Next != nil ; {
				if n2Tmp == n2.Val {
					continue
				}else{
					break
				}
			}
			n2 = n2Tmp
			n1.Next = n2
		}

	}

	return headTmp.Next
}


func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    dummy := &ListNode{0, head}

    cur := dummy
    for cur.Next != nil && cur.Next.Next != nil {
        if cur.Next.Val == cur.Next.Next.Val {
            x := cur.Next.Val
            for cur.Next != nil && cur.Next.Val == x {
                cur.Next = cur.Next.Next
            }
        } else {
            cur = cur.Next
        }
    }

    return dummy.Next
}

func deleteDuplicatesV2(head *ListNode) *ListNode {

	tmp := head

	for ;tmp != nil; {
		if tmp.Next != nil && tmp.Val == tmp.Next.Val {


		}

	}

}

func main() {

}