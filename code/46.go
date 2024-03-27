package main

import (
	"fmt"
)
func permute(nums []int) [][]int {

	var ret [][]int
	var track []int 
	trackMap := make(map[int]bool)

	ret = backtrack(nums, track,trackMap, ret)

	return ret

}

func backtrack(nums[] int,track []int,trackMap map[int]bool,res [][]int) [][]int {
	if len(nums) == len(track) {
		//  避免指针卵用
		res = append(res, append([]int{},track...))
		return res
	}

	for _,item := range nums {
		if v,ok:=trackMap[item];ok && v {
			continue
		}

		trackMap[item] = true

		// var valsTmp = make([]int, len(track))
		// copy(valsTmp,track)
		track = append(track,item)
		res = backtrack(nums,track,trackMap,res)
		// res = append(res, ...)
		trackMap[item] = false
		track = track[:len(track)-1]
	}
    return res
}


func permute2(nums []int) [][]int {
    var res [][]int
    var data []int
    used := make(map[int]int)
    var blacklist func(nums []int)
    blacklist = func(nums []int) {
        if len(data) == len(nums) {
            res = append(res,append([]int{},data...))
            return
        }
        for _,v := range nums {
            if used[v] > 0 {
                continue
            }
            used[v]++
            data = append(data,v)
            blacklist(nums)
            data = data[:len(data )-1]
            used[v]--
        }
    }
    blacklist(nums)
    return res
}

func main() {
	fmt.Println("---permute--",permute([]int{5,4,6,2}))
}