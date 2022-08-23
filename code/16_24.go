package main

import (
	"fmt"
)

func pairSums(nums []int, target int) [][]int {
	reuslt := make([][]int,0,0)

	tmp := make(map[int]int)
	for _,item := range nums {
		tt := target - item

		if v,ok:=tmp[tt];ok && v>0 {
			tmp[tt] -= 1
			reuslt = append(reuslt,[]int{tt,item})
		}else{
			if _,ok:=tmp[item];ok {
				tmp[item] += 1	
			}else{
				tmp[item] = 1
			}
		}

	}

	return reuslt


}

func pairSums_better(nums []int, target int) [][]int {
    m := make(map[int]int)
    ret := make([][]int,0)
    for _, v := range nums {
        t := target - v
        if m[t] >= 1 {
            ret = append(ret, []int{t, v})
            m[t] = m[t] - 1
        } else {
            m[v] = m[v] + 1
        }
    }
    return ret
}

func main() {

	a := make(map[int]int)
	fmt.Println("=======================",a[1],a[2],a[1]+1)

}