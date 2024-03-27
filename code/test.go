package main

import (
	"fmt"
)


func QuickSort(arr []int,left,right int) []int {
	
	if left < right {
		index:=partition(arr,left,right)
		QuickSort(arr,left,index-1)
		QuickSort(arr,index+1,right)
	}

	// index := Index(arr,left,right)
	// if left<index {
	// 	QuickSort(arr,left,index)
	// }
	// if index < right {
	// 	QuickSort(arr,index,index)
	// }


	return arr
}

func partition(arr []int, startIndex, endIndex int) int {
    var (
        pivot = arr[startIndex]
        left  = startIndex
        right = endIndex
    )

    for left != right {                                  
        for left < right && pivot < arr[right] {
            right--
        }

        for left < right && pivot >= arr[left] {
            left++
        }

        if left < right {
            arr[left], arr[right] = arr[right], arr[left]
        }
    }

    arr[startIndex], arr[left] = arr[left], arr[startIndex]

    return left
}



func Index(arr []int,left,right int) int {
	target := arr[(left + right) / 2]
	
	for left <= right {
		for arr[left]<target {
			left++
		}

		for arr[right]>target{
			right-- 
		}

		if left <= right{
			Swap(arr,left,right)
			left ++ 
			right -- 	
		}
		
		
	}
	return left
}

func Swap(arr []int,i,j int) []int{
	arr[i],arr[j] = arr[j],arr[i]
	return arr
}


func Find(arr []int,target int) int {
	left,right := 0,len(arr)-1
	for left <= right {
		mid := (left+right)/2
		if target == arr[mid] {
			return mid
		}else if arr[mid] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}

	return -1
}

func B(a string) []string {
	ret := []string{}
	if len(a) == 1 {
		return []string{a}
	}

	tmp := B(a[1:])

	for _,item := range tmp {
		for i:=0;i<len(item);i++{
			ret = append(ret, item[0:i] + string(a[0]) + item[i:])
		}

		ret = append(ret, item + string(a[0]))
	}
	fmt.Println("-------------",a,ret)
	return ret
}

func main() {

	fmt.Println("----1----",B("123"))
	a := []int{1,2,3,4,5}
	fmt.Println("----------------------",Find(a,2),Swap(a,1,2))
}



