package main

import (
	"fmt"
)

func abs(a,b int)int{
	ret := a-b
	if ret >= 0 {
		return ret
	}
	return 0-ret
}

func Min(a,b int) int {
	if a < b {
		return a
	}else{
		return b
	}
}


func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	min_total := 0

	for _,item := range  baseCosts {
		
		// min_tmp := abs(total,target)
		// if min_tmp < min {
		// 	min = min_tmp
		// }
		min_total = closest(toppingCosts, 2,item, target))

	}

	return target 
}


// min 
func closest(toppingCosts []int,count,closed, target int) int {

	if abs(closed,target) {

	}

	// Min(min,target)

	if count == 0 {
		return min
	}

	for _,item := range toppingCosts {

		closest()


	}




	return 
}


func main() {
	fmt.Println("=======================")
}
