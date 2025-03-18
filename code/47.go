package main

import (
	"fmt"
)



func permuteUnique(nums []int) [][]int {
	var ret [][]int
	var track []int

	ret = backtrack2(nums,track,ret)

	return ret
}

func backtrack2(nums []int,track []int,res [][]int) [][]int {
	if len(nums) == 0 {
		res = append(res, append([]int{},track...))
		return res
	}

	for index,item := range nums {
		track = append(track, item)
		tmp := []int{}
		tmp = append(tmp,nums[:index]...)
		tmp = append(tmp,nums[index+1:]...)
		fmt.Println("---------s------",index,item,tmp)
		res = backtrack2(tmp,track,res)

		fmt.Println("---------e------",index,item,res)
		track = track[:len(track)-1]
	}
	return res
}


func main() {
	fmt.Println("---permute--",permuteUnique([]int{1,2,3}))
}