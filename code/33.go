package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left,right := 0,len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		//  后半截是有序的
		if nums[mid] < nums[right] {
			if target > nums[mid] &&  target <= nums[right] {
				left = mid + 1
			}else{
				right = mid -1 
			}

		}else{
			//  前半截是有序的
			if target  < nums[mid] && target >= nums[left] {
				right = mid - 1
			}else{
				left = mid + 1
			}

		}

	}


	return -1
}

func main() {
	fmt.Println("------------",search([]int{1,3},3))
}