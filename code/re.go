package main


import (
	"fmt"
)


type Node struct {
	LeftChild  *Node 
	RightChild *Node
	Val int
}


func Range(t *Node) {
	if t == nil {
		return 
	}

	fmt.Println("---",t.Val)

	if t.LeftChild != nil {
		Range(t.LeftChild)
	}

	if t.RightChild != nil {
		Range(t.RightChild)
	}
}


func main() {
	t:= &Node{Val:1}
	t.LeftChild = &Node{Val:2}
	t.RightChild = &Node{Val:3}

	t.LeftChild.RightChild = &Node{Val:4}

	t.RightChild.LeftChild = &Node{Val:5}
	t.RightChild.RightChild = &Node{Val:6}

	Range(t)
}


// 

//		1

//	2 		3
//		4 5		6
