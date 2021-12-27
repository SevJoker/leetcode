package main

import (
	"fmt"
)

func numberOfSubarrays(nums []int, k int) int {
	IndexMap := map[int]int{}
	indexSlice := []int{}
	i := 0
	result := 0
	for index,item := range nums {
		if item%2 != 0 {
			IndexMap[index] = i
			i = 0
			indexSlice = append(indexSlice,index)
			if len(indexSlice) > k {
				start := IndexMap[indexSlice[0]]
				end := IndexMap[index]
				result += (start+1) * (end+1)
				indexSlice = indexSlice[1:]
			}
		}else{
			i += 1	
		}
	}
	fmt.Println("===============",indexSlice,IndexMap,i)

	if len(indexSlice) == k {
		start := IndexMap[indexSlice[0]]
		end := i
		result += (start+1) * (end+1)
	}

	return result


}

func numberOfSubarrays1(nums []int, k int) int {
	var ans int
	n := len(nums)
	s := make([]int, n + 1)
	s[0] = 0
	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1] % 2
	}

	//iMap := make(map[int]int)
	count := make([]int, n + 1)
	count[s[0]]++
	for i := 1; i <= len(nums); i++ {
		if s[i] - k >= 0 {
			ans = ans + count[s[i] - k]
		}
		count[s[i]]++
	}

	return ans
}

func numberOfSubarrays2(nums []int, k int) int {
	arr := make([]int,len(nums)+1)
	arr[0] =1 
	sum,res := 0,0
	for _,item := range nums {
		sum += item%2
		arr[sum] ++ 
		if sum >= k {
			res += arr[sum-k]
		}
	}

	return res
}


func main() {
	fmt.Println("--------main-------",numberOfSubarrays2([]int{2,2,2,1,2,2,1,2,2,2},2))
}
