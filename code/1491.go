package main

func average(salary []int) float64 {
	if len(salary) < 3 {
		return 0
	}

	min,max,sum := salary[0],salary[0],0
	for _,item := range salary  {
		if min > item {
			min = item
		}
		if max < item {
			max = item
		}
		sum += item
	}

	sum = sum - min - max
	return float64(sum)/float64(len(salary)-2)
}


func main() {

}