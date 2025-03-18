package main

import (
	"fmt"
)


type LRUCache1 struct {
	HashMap  map[int]*Node
	Cache    *DoubleList
	Cap      int // 最大容量
}

type Node struct {
	val 	int 
	key 	int
	prev 	*Node
	next 	*Node
}

func NewNode(k,v int) *Node {
	return &Node{
		val:v,
		key:k,
	}
}

type DoubleList struct {
	head  *Node
	tail  *Node
	size  int // 个数
}

func NewDoubleList() *DoubleList {
	ret := &DoubleList{
		head: NewNode(0,0),
		tail: NewNode(0,0),
		size: 0,
	}

	ret.head.prev = ret.tail
	ret.tail.next = ret.head
	return ret
}

func (d *DoubleList) Print() {
	n := d.tail
	for ;n!=nil;n=n.next {
		fmt.Println("---DoubleList-Print-",n.key,n.val)
	}
}


func (d *DoubleList) Add(x *Node) {
	x.prev = d.head.prev
	x.next = d.head
	d.head.prev.next = x
	d.head.prev = x
	d.size ++
}

func (d *DoubleList) Remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	d.size --
}


func (d *DoubleList) RemoveLast() *Node {
	if d.tail.next == d.head {
		return nil
	}

	n := d.tail.next
	d.Remove(n)
	return n
}



func Constructor(capacity int) LRUCache1 {
	ret := LRUCache1{
		HashMap: make(map[int]*Node),
		Cap :capacity,
		Cache : NewDoubleList(),
	}
	return ret
}


func (this *LRUCache1) putFirst(key int) {
	n,ok := this.HashMap[key]
	if !ok {
		return 
	}

	this.Cache.Remove(n)
	this.Cache.Add(n)
}

func (this *LRUCache1) add(key,value int) {
	n := NewNode(key,value)
	this.Cache.Add(n)
	this.HashMap[key] = n
}

func (this *LRUCache1) remove(key int) {
	n,ok := this.HashMap[key]
	if !ok {
		return
	}
	this.Cache.Remove(n)
	delete(this.HashMap,key)
}
// 删除最老数据
func (this *LRUCache1) removeold() {
	n := this.Cache.RemoveLast()
	if n==nil {
		return
	}
	delete(this.HashMap, n.key)
}


func (this *LRUCache1) Get(key int) int {
	n,ok:= this.HashMap[key]
	if ok {
		this.putFirst(key)
		return n.val
	}
	return -1
}


func (this *LRUCache1) Put(key int, value int)  {
	_,ok:= this.HashMap[key]
	if ok {
		this.remove(key)
		this.add(key,value)
		return 
	}

	if this.Cap <= this.Cache.size {
		// 删除最老的
		this.removeold()
	}

	this.add(key, value)
}



// 顺时针打印二维数据

func spiralArray(array [][]int) []int {	
	if len(array) == 0 {
		return []int{}
	}
	rows,columns := len(array),len(array[0])
	ret := make([]int,rows*columns)
	index := 0
	left,right,top,buttom := 0,columns-1,0,rows-1

	for left <= right && top <= buttom {
		for column := left;column <= right; column ++ {
			ret[index] = array[top][column]
			index ++ 

		}
		for row := top+1; row <= buttom;row ++ {
			ret[index] = array[row][right]
			index ++ 
		}
		if left<right && top < buttom{
			for column := right-1;column>left;column--{
				ret[index] = array[buttom][column]
				index++
			}

			for row := buttom;row>top;row--{
				ret[index] = array[row][right]
				index++
			}
		}
		left ++
		right--
		top ++
		buttom --
	}
	
	return ret

}



func main() {

	// lru:=Constructor(3)

	// lru.Put(1,1)
	// lru.Put(2,2)
	// lru.Put(3,3)
	// lru.Cache.Print()
	// lru.Put(4,4)
	// lru.Cache.Print()

	fmt.Println("-----------------")

	arr := [][]int{[]int{1,2,3,4,5},[]int{1,2,3,4,5},[]int{1,2,3,4,5}}
	// []int{1,2,3,4,5},[]int{1,2,3,4,5},[]int{1,2,3,4,5}
	// arr = append(arr,[]int{1,2,3,4,5})
	// arr = append(arr,[]int{1,2,3,4,5})
	// arr = append(arr,[]int{1,2,3,4,5})

	fmt.Println("======spiralArray=====",spiralArray(arr))
}





// 更简单

// 添加头尾节点指针，并且在初始化的时候把它们连接起来，可以大大简化判断，代码更简洁
// 因为访问时涉及到挪动操作，要在 o(1) 的时间复杂度实现，只能用链表
// type LRUCache struct {
//     m map[int]*Node
//     capacity int
//     count int
//     // 头尾指针不存数据，单向链接到头尾节点
//     head *Node
//     tail *Node
// }

// type Node struct {
//     key int
//     val int
//     prev *Node
//     next *Node
// }

// func Constructor(capacity int) LRUCache {
//     l := LRUCache{
//         m: make(map[int]*Node, capacity),
//         capacity: capacity,
//         head:&Node{},
//         tail:&Node{},
//     }
//     l.head.next = l.tail
//     l.tail.prev = l.head
//     return l
// }


// func (this *LRUCache) Get(key int) int {
//     n, ok := this.m[key]
//     if !ok {
//         return -1
//     }
//     this.RemoveNode(n)
//     this.AddToHead(n)
//     return n.val
// }


// func (this *LRUCache) Put(key int, value int)  {
//     n, ok := this.m[key]
//     if ok {
//         n.val = value
//         this.RemoveNode(n)
//         this.AddToHead(n)
//         return
//     }

//     if this.count == this.capacity {
//         tail := this.RemoveTail()
//         delete(this.m, tail.key)
//         this.count--
//     }
//     this.count++
//     n = &Node{key:key, val:value}
//     this.AddToHead(n)
//     this.m[key] = n
// }

// func (this *LRUCache) RemoveNode(n *Node) {
//     n.prev.next = n.next
//     n.next.prev = n.prev
//     n.next = nil
//     n.prev = nil
// }

// func (this *LRUCache) RemoveTail() *Node {
//     tail := this.tail.prev
//     this.RemoveNode(tail)
//     return tail
// }

// func (this *LRUCache) AddToHead(n *Node) {
//     // 改变原先头结点的前驱节点为新的节点
//     n.next = this.head.next
//     this.head.next.prev = n
//     // 改变新头结点的前驱节点为 head 指针
//     this.head.next = n
//     n.prev = this.head
// }


// func init() { debug.SetGCPercent(-1) }