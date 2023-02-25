package main

import (
	"fmt"
)


func longestConsecutive(nums []int) int {
	nums_map := make(map[int]int)
	max := 0
	for _,item := range nums {
		if _,ok:=nums_map[item];ok{
			continue
		}

		l := length(nums_map,item)
		if l > max {
			max = l
		}
		union(nums_map,item)
	}

	return max	
}


func length(nums_map map[int]int,target int) int {
	l := 1
	if v,ok:=nums_map[target-1];ok {
		l += v
	}

	if v,ok :=nums_map[target+1];ok{
		l += v
	}

	nums_map[target] = l
	return l
}

func union(nums_map map[int]int,target int) {
	for left:=target-1;;{
		if _,ok:=nums_map[left];!ok{
			break
		}
		nums_map[left] = nums_map[target]
		left=left-1
	}
	for right:=target+1;;{
		if _,ok:=nums_map[right];!ok{
			break
		}
		nums_map[right] = nums_map[target]
		right=right+1
	}
}

func main() {

	fmt.Println("-----------------",longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1}))
}	