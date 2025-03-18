package main

import (
	"fmt"
)

/*
	核心奥义
		1. 确认状态
		2. 确认dp函数定义
		3. 明确选择
		4. 明确 bad case： 起始值处理


	标准范式
		无范式，核心根据奥义对照处理

		func dpfunc(args) {
			dp := xxx  初始化dp

			for i in 状态值 {


				for j in 选择列表 {
					做选择，有些简单的不需要做选择 直接推导即可

					dp[i] = func(dp[i-n]) 根据转换函数处理dp值
				}

			}


			return func2(dp[i])
		}

*/



//  斐波拉契算法：动态规划解法
//   dp[n] = dp[n-1] + dp[n-2]
func fibdp(n int) int {
	dp := make([]int,n)
	dp[0] = 1
	dp[1] = 1
	for i:=2;i<n;i++{
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n-1]
}



// 求最少能凑成amount的金额数
//   状态   n = amount
//   dp   dp[n] = dp[n-coin] + 1  
//   选择  lists中遍历选择
//   bad case   n = 0 dp[n] = 0  n<-1  dp[n] = -1（无解）
func coinChange(lists []int,amount int) int {
	dp := make([]int,amount+1)
	for i,_ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i:=1;i<=amount;i++ {
		for _,coin := range lists {
			if i - coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
			fmt.Println("------------i-----",i,dp)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}


func min(x,y int) int {
	if x<y {
		return x
	}else{
		return y
	}
}



func main() {

}