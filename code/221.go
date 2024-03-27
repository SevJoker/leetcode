package main

import (
	"fmt"
)




func maximalSquare2(matrix [][]byte) int {
	dp := make([][]int,len(matrix))
	max := 0
	
	length := len(matrix)
	hight := len(matrix[0])
	for i:=0;i<length;i++ {
		dp[i]= make([]int,hight)

		for j:=0;j<hight;j++{
			tmp:= int(matrix[i][j] - '0')
			if tmp == 0 {
				dp[i][j] = tmp
			}else{
				if i > 0 && j > 0{
					dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1]) + 1 
				}else{
					dp[i][j] = tmp
				}
			}


			if max < dp[i][j] {
				max = dp[i][j]
			}
		}


	}
	return max*max
}

func min(args ...int) int {
	min := args[0]
	for _,i := range args{

		if min > i {
			min = i
		}
	}
	
	return min
}

func maximalSquare(matrix [][]byte) int {
	length := len(matrix)
	hight := len(matrix[0])

	max := 0
	for i:=0;i<length;i ++ {
		for j:=0;j<hight;j++ {
			if string(matrix[i][j]) == "1" {
				tmp := searchMax(matrix,i,j)
				if max < tmp {
					max = tmp
				}
			}

            if max > hight - i  {
                break
            }    
		}
        if max > length - i  {
            break
        }
	}
	return max * max
}

func searchMax(matrix [][]byte,i,j int) int {
	length := len(matrix)
	hight := len(matrix[0])
	ret := 0
	for ii:=0;;ii++ {
		if ii+i >= length {
			break
		}
		if ii+j >= hight {
			break
		}

		flag := false

		for jj:=0;jj <= ii;jj++ {
			if string(matrix[i+ii][j+jj]) !="1" {
				flag = true
				break
			}
			if string(matrix[i+jj][j+ii]) !="1" {
				flag = true
				break
			}

			
		}

		if flag {
			break
		}
		ret = (ii+1)
	}
    return ret
}

func main() {

}