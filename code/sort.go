package main

import (
	"fmt"
)


func QuickSort(lists []int) []int {

	return quickSort(lists,0,len(lists)-1)
}

func quickSort(lists []int,low,high int) []int{
	if low < high {
		pivot := partition(lists,low,high)
		quickSort(lists,low,pivot-1)
		quickSort(lists,pivot+1,high)
	}
	return lists
}


func partition(lists []int,low,high int) int {
	// 直接取第一个作为锚点
	pivot := lists[low]
	//  不断移动位置 找到  pivot  的坐标并返回
	for low < high {
		for ;low < high && lists[high] >= pivot;{
			high --
		}
		lists[low] = lists[high]
		for ;low < high && lists[low] <= pivot;{
			low ++
		}
		lists[high] = lists[low]
	}
	lists[low] = pivot
	return low
}


func main() {
	fmt.Println("-------------------main---------------",QuickSort([]int{2,1,10,1,2,3,3,11}))
	// fmt.Println("-------------------main---------------",QuickSort([]int{2,1,10,3,11}))
}