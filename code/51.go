package main

import (
	"fmt"
	"strings"
)


func solveNQueens(n int) [][]string {

	var res [][]string

	var track  = make([][]bool,n)

	for index := range track {
		track[index] = make([]bool,n)
	}

	var (
        bracktrack func(n int,row int)
    )

	validmap := make(map[int]bool)

	bracktrack = func(n int,row int) {
		if row == n {
			fmt.Println("--------",track,res)
			tmp := GenQueens(track)
            res = append(res, tmp)
			return 
		}

		for i:=0;i<n;i++ {
			//  不合法判断
			if v,ok := validmap[i];(v && ok) || isValid(track,row,i){
				continue
			}
			//  如果斜对角也不对

			track[row][i] = true
			validmap[i] = true

			bracktrack(n,row + 1)
			//  回撤
			track[row][i] = false
			validmap[i] = false
		}


	}

    bracktrack(n,0)

	return res
}


func isValid(track [][]bool,row,col int) bool {
	//  斜对角判断
	j:=col+1
	for i:=row-1;i>=0&&j<len(track);i-- {
		if track[i][j] {
			return true
		}
		j++
	}

	j=col-1
	for i:=row-1;i>=0&&j>=0;i-- {
		if track[i][j] {
			return true
		}
		j--
	}

	return false
}


func GenQueens(track [][]bool) []string {
	ret := []string{}
	
	for _,item :=range  track {

		// tmp:= ""

		var builder strings.Builder

		for _,flag := range item {
			if flag {

				builder.WriteString("Q")
				// tmp = tmp + "@"
			}else{
				builder.WriteString(".")
				// tmp = tmp + "."
			}
		}

		ret = append(ret,builder.String())
	}


	return ret
}



func testString() {
	a := "aaa"
	// a[1] = "b"

	var str string


	fmt.Println("-------a---",a,str)
}

func main() {
	testString()
	fmt.Println("---permute--",solveNQueens(8))
}