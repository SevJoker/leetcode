package main

import (
	"fmt"
)

/*
	核心奥义
		1. 二分法
		2. 滑动窗口
			- 双指针
		3. 回文串
		4. 前置和  技巧（解决子数组问题）
			- 核心思路就是 缓存另外一个数组来记录 0-i的和
		5. 查分数组（解决批量修改的问题）


*/




//  二分查找
func find2(nums []int,target int) int {
	left,right := 0,len(nums)-1
	mid := 0

	for left <= right {
		mid = left + (right-left)/2
		if target == nums[mid] {
			return mid
		}else if target > nums[mid] {
			left = mid + 1
		}else{
			right = mid -1
		}

	}



	return -1
}


// 快速排序
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

}