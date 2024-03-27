package main

import (
	"fmt"
)


func permutation(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	var ret map[string]bool
	for _,item := range permutation(s[1:]) {
		for i:=0;i<=len(item);i++ {
			ret[item[0:i]+string(s[0]) + item[i:]] = true
		}
	}

	var retSlice []string
	for k,_:=range ret {
		retSlice = append(retSlice,k)
	}

	return retSlice
}

func main() {
	fmt.Println("--------------",permutation("abc"))
}


func permutation(s string) (ans[]string) {
    loop:=[]byte(s)
    sort.Slice(loop,func(i,j int)bool{return loop[i]<loop[j]})
    for{
        ans = append(ans,string(loop))
        if !nextPermutation(loop){
            break
        }
    }
    return
}
func reverse(nums []byte){
    for i,j:=0,len(nums)-1;i<len(nums)/2; i,j=i+1,j-1{
        nums[i],nums[j] = nums[j],nums[i]
    }
}
func nextPermutation(nums []byte) bool{
    //1.find k
    k:=-1
    for i:=0; i < len(nums)-1; i=i+1{
        if nums[i] < nums[i+1]{
            k = i
        }
    }
    if k == -1{
        reverse(nums)
        return false
    }
    //2. find l
    l:=k+1
    for i:=k+1; i<len(nums); i=i+1{
        if(nums[i]>nums[k]){
            l = i
        }
    }
    nums[l],nums[k] = nums[k],nums[l]
    reverse(nums[k+1:])
    return true
}




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