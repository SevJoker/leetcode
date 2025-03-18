package main

import (
	"fmt"
)


type ListNode struct {
    Val int
 	Next *ListNode
 }


func QuickSort(arr []int,left,right int) []int {
	
	if left < right {
		index:=partition(arr,left,right)
		QuickSort(arr,left,index-1)
		QuickSort(arr,index+1,right)
	}

	// index := Index(arr,left,right)
	// if left<index {
	// 	QuickSort(arr,left,index)
	// }
	// if index < right {
	// 	QuickSort(arr,index,index)
	// }


	return arr
}

func partition(arr []int, startIndex, endIndex int) int {
    var (
        pivot = arr[startIndex]
        left  = startIndex
        right = endIndex
    )

    for left != right {                                  
        for left < right && pivot < arr[right] {
            right--
        }

        for left < right && pivot >= arr[left] {
            left++
        }

        if left < right {
            arr[left], arr[right] = arr[right], arr[left]
        }
    }

    arr[startIndex], arr[left] = arr[left], arr[startIndex]

    return left
}



func Index(arr []int,left,right int) int {
	target := arr[(left + right) / 2]
	
	for left <= right {
		for arr[left]<target {
			left++
		}

		for arr[right]>target{
			right-- 
		}

		if left <= right{
			Swap(arr,left,right)
			left ++ 
			right -- 	
		}
		
		
	}
	return left
}

func Swap(arr []int,i,j int) []int{
	arr[i],arr[j] = arr[j],arr[i]
	return arr
}


func Find(arr []int,target int) int {
	left,right := 0,len(arr)-1
	for left <= right {
		mid := (left+right)/2
		if target == arr[mid] {
			return mid
		}else if arr[mid] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}

	return -1
}

func B(a string) []string {
	ret := []string{}
	if len(a) == 1 {
		return []string{a}
	}

	tmp := B(a[1:])

	for _,item := range tmp {
		for i:=0;i<len(item);i++{
			ret = append(ret, item[0:i] + string(a[0]) + item[i:])
		}

		ret = append(ret, item + string(a[0]))
	}
	fmt.Println("-------------",a,ret)
	return ret
}


func Print() {

}

//  微服务 a，b  是否有循环调用

//  dfs遍历即可
func FindWei() {
	calls := [][2]string{
        {"A", "B"},
        {"B", "C"},
        {"C", "D"},
        {"D", "C"},
    }

    g := make(map[string][]string)
    for _,item := range calls {
    	if _,ok:=g[item[0]];ok {
    		g[item[0]] = append(g[item[0]],item[1])
    	}else{
    		g[item[0]] = []string{item[1]}
    	}
    }


    fmt.Println("-------Find------",Cycle(g))
}

func Cycle(g map[string][]string) bool {
	visited := make(map[string]bool)
	// visitedtmp := make(map[string]bool)
	for k,_ := range g {
		if !visited[k] {
			visited[k] = true
			visitedtmp := make(map[string]bool)
			if hasCycle(g,k,visitedtmp) {
				return true
			}
			
		}

	}
	return false
}

func hasCycle(g map[string][]string,node string,visitedtmp map[string]bool) bool {
	visitedtmp[node] = true
	v,ok := g[node]
	if !ok {
		return false
	}

	for _,item := range v {
		if visitedtmp[item] {
			return true
		}
		return hasCycle(g,item,visitedtmp)
	}


	return false
}


//  206 反转链表
func reverse(head *ListNode) *ListNode {
	// vh := &ListNode{Next:head}

	var (
		cur = head
		pre *ListNode = nil
	)

	for cur != nil {
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
		fmt.Println("---ListNode----",r)
		r = r.Next	
	}

	fmt.Println("----1----",B("123"))
	a := []int{1,2,3,4,5}
	fmt.Println("----------------------",Find(a,2),Swap(a,1,2))
	FindWei()
}



