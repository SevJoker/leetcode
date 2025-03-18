package main

import (
	"fmt"
	"sort"
	"strings"
)


/*
	核心奥义：
		1. 路径：记录已经做过的选择
		2. 选择列表：当前可以做的选择
			- 选择完记得回撤
		3. 结束条件：决策树的末尾处理

	标准范式

		func huisu(args) {

			ret := backtrack(args)
			return ret
		}

		func backtrack(args) {
			if 决策树末尾 {
				return 
			}

			for item in 选择列表 {
				判断选择是否有效

				dosth 做选择

				res = backtrack(args)

				dosth 回撤选择
			}
		}
*/

//  给一个数字  23212 n  一个数组 A = {1,2,3} 用A里面的数字拼接成小于n的最大值
func maxInt(n int,A []int ) int {
	na := []int{}
	tmp := n
	for i:=0; tmp>0; i++ {
		na = append(na, tmp%10)
		tmp = tmp/10
	}

	ret := -1
	
	var track = []int{}

	var (
        bracktrack func(i int) int
    )

	bracktrack = func(i int) int {
		if len(track) == len(na) {
			tmp := ToInt(track)
			if tmp >= n {
				return -1
			}

			if tmp > ret {
				ret = tmp
			}

			return tmp
		}
		for _,item := range A {
			//  可以避免无效选择 加个判断


			track = append(track, item)
			bracktrack(i-1)
			track = track[:len(track)-1]
		}
		return -1
	}
	bracktrack(len(na)-1)


	// 看是否能降位处理
    if ret == -1 && len(na) > 1 {
		// 生成降位之后的最大值即可
		maxA := max(A)
		ret = 0
		tmp = 1
		for i:=0;i<len(na)-1;i++ {
			ret = ret + maxA*tmp
			tmp = tmp * 10
		}
    }

	return ret
}

func max(l []int) int  {
	max := -1 
	for _,item := range l {
		if item > max {
			max = item	
		}

	}
	return max
}

func ToInt(na []int) int {
	ret := 0
	tmp := 1
	for _,item := range na {
		ret = ret + item * tmp
		tmp = tmp * 10
	}
	return ret
}



//  全排列问题  lc  47

func permuteUnique(nums []int) [][]int {
	var ret [][]int
	var track []int

	ret = backtrack2(nums,track,ret)

	return ret
}

func backtrack2(nums []int,track []int,res [][]int) [][]int {
	if len(nums) == 0 {
		res = append(res, append([]int{},track...))
		return res
	}

	for index,item := nums {
		track = append(track, item)
		tmp := []int{}
		tmp = append(tmp,nums[:index]...)
		tmp = append(tmp,nums[index+1:]...)
		res = backtrack2(nums,tmp,res)
		track = track[:len(track)-1]
	}
	return res
}

//  全排列问题  lc  46
//   定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	var ret [][]int
	var track []int 
	trackMap := make(map[int]bool)
	ret = backtrack(nums, track,trackMap, ret)
	return ret

}

func backtrack(nums[] int,track []int,trackMap map[int]bool,res [][]int) [][]int {
	if len(nums) == len(track) {
		//  避免指针卵用
		res = append(res, append([]int{},track...))
		return res
	}

	for _,item := range nums {
		if v,ok:=trackMap[item];ok && v {
			continue
		}

		trackMap[item] = true

		// var valsTmp = make([]int, len(track))
		// copy(valsTmp,track)
		track = append(track,item)
		res = backtrack(nums,track,trackMap,res)
		// res = append(res, ...)
		trackMap[item] = false
		track = track[:len(track)-1]
	}
	return res
}


//  N皇后问题
//   保证横竖斜 都不能有重叠
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

//  

/*
思路：
1.回溯：也就是DFS，为了防止字符串内有重复字符，可以先将字符串排序，然后每次取字符串，都保证从相同字符的最左边一个取起。
2.下一个排列：由于字符有大小关系，可以先生成字符串的最小排列，然后快速得到字典序中下一个更大的排列。
*/
func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(nums []byte) bool {
	n := len(nums)
	i := n - 2	//从最后开始往前找，找到第一个值比右边小的元素i
	for i >= 0 && nums[i] >= nums[i+1] {	
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for j >= 0 && nums[i] >= nums[j] {	//从右往左找到第一个比此元素i大的元素j
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]	//交换i，j位置。
	reverse(nums[i+1:])	//逆向所有原本i位置右边的数（将逆向变成了顺向）
	return true
}

func permutation(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for {
		ans = append(ans, string(t))
		if !nextPermutation(t) {
			break
		}
	}
	return
}


func main() {

	fmt.Println("--------maxInt------",maxInt(5555,[]int{5,9}))

}