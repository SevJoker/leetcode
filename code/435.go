package main

import (
	"fmt"
)



func eraseOverlapIntervals1(intervals [][]int) int {
     n := len(intervals)
    if n == 0 {
        return 0
    }
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if intervals[j][1] <= intervals[i][0] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }
    return n - max(dp...)
}

func max(a ...int) int {
    res := a[0]
    for _, v := range a[1:] {
        if v > res {
            res = v
        }
    }
    return res
}




type A [][]int

func (a A) Less(i, j int) bool {
	return a[i][1] > a[j][1]
}

func (a A) Swap(i,j int) {a[i],a[j]=a[j],a[i]}


func main() {

	fmt.Println("-----------------",eraseOverlapIntervals([]int{[1,2],[2,3],[3,4],[1,3]}))
}	