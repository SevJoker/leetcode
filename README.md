# leetcode梳理




### 递归解法
-思路 根据题目定义一个可递归的计算公式
eg 236 ： 左右均找到即为 root，不然就取左 或 右

### 前缀和算法
-说明 预处理，生成一个新数组，当前下标代表原始数据改下标之前所有的累计值 即为b[i] = b[i-1]+a[i]

