package main

import (
	"fmt"
	"math"
)

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


const inf = math.MaxInt32
func coinChange2(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i :=1; i<=amount; i++ {
        dp[i] = inf
    }
    for i := range dp {
        for _, c:= range coins {
            if i-c >= 0 {
                dp[i] = min(dp[i], dp[i-c]+1)
            }
        }
    }
    if dp[amount] == inf {
        return -1
    }
    return dp[amount]
}

func main() {

	fmt.Println("----- coinChange()-------", coinChange([]int{2},3))
}