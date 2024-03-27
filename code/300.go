package main

import (
	"fmt"
)

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

func lengthOfLIS(nums []int) int {

	dp := []int{}
	max := 1

	for index,item := range nums {
		dp = append(dp,1)

		for j:=0;j<index;j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i],dp[j]+1)
			}
		}

		max = Max(dp[i],max)
	}


	return max

}



//  动态规划思路 
//     核心要找到动态规划的公式  即为  dp[i] = xxx(dp[j]) 可以有其他限制 可以定义dp[i]的属性，保证推演过程的研究
//   

//  这一题而言
//     dp[i]为选取nums[i]的情况下前面的最长子序列
// 	   推演出 dp[i] = max(dp[j]) + 1 (j<i 且nums[j]<nums[i])
//     最后 dp数组中的最大值即为答案
func Max(args ...int) int {
	m := args[0]
	for item := range args {
		if m < item {
			m = args
		}
	}
	return m
}

func main() {

}