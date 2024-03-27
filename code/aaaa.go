package main	


import (
	"fmt"
)


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

//  快排



//  https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
//  开胃小菜 合并两个月有序数组
func mergebad(nums1 []int, m int, nums2 []int, n int) []int {
    var i,j = 0,0
    var ret []int
    for index:=0;index < m+n;index++ {
        if i < m && j < n {
            if nums1[i] <= nums2[j] {
			    ret = append(ret,nums1[i])
			    i++
            } else {
				ret = append(ret,nums2[j])            	
				j++	
            }
            
        } else if i < m {
        	ret = append(ret, nums1[i])
        	i++
        } else if j < n {

        	ret = append(ret, nums2[j])
        	j++
        }
    }

    return ret
}





func merge(nums1 []int, m int, nums2 []int, n int) []int {
	p1 := m-1
	p2 := n-1
	p := m+n - 1

	for ;p1 >=0 && p2>=0 ;{
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1 -- 
		}else{
			nums1[p] = nums2[p2]
			p2 -- 
		}
		p --
	}

	for ;p1 >=0; {
		nums1[p] = nums1[p1]		
		p1 -- 
		p--
	}
	for ;p2>=0; {
		nums1[p] = nums2[p2]
		p2 -- 
		p--
	}

	return nums1
}








// https://leetcode.cn/problems/generate-parentheses/?envType=study-plan-v2&envId=top-interview-150
//   数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 输入：n = 1
// 输出：["()"]
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]

//  核心思路  F[n] = "("+F[i]+")” + F(n-i-1)
func generateParenthesis2(n int) []string {
    dp := [][]string{}
	dp = append(dp,[]string{""})
	dp = append(dp,[]string{"()"})
	for i:=2;i<= n ;i++ {
		l := []string{}
		for j:=0;j<i;j++ {
			for _,itemj := range dp[j] {
				for _,itemk := range dp[i-j-1] {
					l = append(l,"("+itemj+")"+itemk)
				}
			}
		}

		dp = append(dp, l)
	}

	return dp[n]
}



func generateParenthesis(n int) []string {
	result := []string{}
	backtrack("", 0, 0, n, &result)
	return result
}

func backtrack(current string, open, close, max int, result *[]string) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(current+"(", open+1, close, max, result)
	}
	if close < open {
		backtrack(current+")", open, close+1, max, result)
	}
}



// 斐波拉契数列
//  递归解法
func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}


//  动态规划解法
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


func main() {

	fmt.Println("-------fibdp-------",fib1(5),fibdp(5))

	fmt.Println("----testmrge---",merge([]int{1,2,3,0,0,0},3,[]int{2,5,6},3))

	fmt.Println("----testmrge---",generateParenthesis(2))

}