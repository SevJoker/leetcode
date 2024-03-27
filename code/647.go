package main

import (
	"fmt"
)


// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

// 回文字符串 是正着读和倒过来读一样的字符串。

// 子字符串 是字符串中的由连续字符组成的一个序列。

// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。



func countSubstringsn3(s string) int {
	ss := []rune(s)
	// for i,w := range s {
	length := len(ss) -1
	a,b,count:=0,0,0
	ishit := false
	for i:=0;i<=length;i++ {

		for si:=length;i<=si;si--{
			a,b = i,si

			for a <= b {
				if ss[a] != ss[b] {
					ishit = false
					break
				} else {
					ishit = true
					a ++ 
					b -- 
				}
			}

			if ishit {
				count ++ 
			}
		}

	}

	return count
}


func countSubstringsn(s string) int {

	ss := []rune(s)
	// for i,w := range s {
	length := len(ss) -1
	count:=0
	dp := make([]map[int]bool,length+1,length+1)
	for i:=0;i<=length;i++ {
		dp[i] = make(map[int]bool)
	}

	for j:=0;j<=length;j++ {

		for i:=j;i>=0;i-- {
			if ss[i]==ss[j] && (j-i<2||dp[i+1][j-1]) {
				count ++
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}

	}

	return count
}


func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][i] = true
		if i+1 < n {
			if s[i] == s[i+1] {
				dp[i][i+1] = true
			}
		} else {
			continue
		}
		for j := i + 2; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] {
				res++
			}
		}
	}

	return res
}


func main() {
	fmt.Println("--------------------",countSubstrings("assdda"))
}