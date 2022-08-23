package main

import (
	"fmt"
	"sort"
)


// 0-8  => 万
// 9-17 => 条
// 18-26 => 条
// 27-30 => 东南西北
// 31-33 => 中发白

func success(data []int) {
	if len(data) != 14 {
		return false
	}


	sort.Ints(data)
	fmt.Println(data)

	

	return false
}

func main() {
	a := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println("---------------",success())
}

