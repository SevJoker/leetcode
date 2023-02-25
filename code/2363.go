package main

import (
	"fmt"
	"sort"
)


type Sorter [][]int

//排序 最新购买的城市/景区优先，同一城市/景区一起展示，同一模块优先月卡，再次卡，次卡间按照最近购买优先展示
func (x Sorter) Len() int      { return len(x) }
func (x Sorter) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x Sorter) Less(i, j int) bool {
	if x[i][0] > x[j][0] {
		return false
	}
	return true
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {


 	indexM := make(map[int]int)
 	resut := make([][]int,0,0)

 	for index,item := range items1 {
 		resut = append(resut,item)
 		indexM[item[0]] = index
 	}

 	for _,item := range items2 {
 		if v,ok:= indexM[item[0]];ok {
 			resut[v][1] += item[1]
 		}else{
 			resut = append(resut,item)
 		}
 	}

 	sort.Sort(Sorter(resut))

    return resut
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := make(map[int]int)
	for _, val := range items1 {
		m[val[0]] += val[1]
	}
	for _, val := range items2 {
		m[val[0]] += val[1]
	}
	ans := make([][]int, 0)
	for k, v := range m {
		ans = append(ans, []int{k, v})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

func main() {
	a := [][]int{{1,1},{4,5},{3,8}}
	b := [][]int{{1,3},{10,10},{3,8}}
	fmt.Println(mergeSimilarItems(a,b))
}