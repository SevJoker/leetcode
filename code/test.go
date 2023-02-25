package main


// 标题
// 删除链表重复元素

// 题目描述
// 存在一个按升序排列的单向链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中【没有重复出现】的数字，返回同样按升序排列的结果链表。​

// 输入  1->1->1->2->3->4->4​
// 输出   2->3​

type ListNode struct {​ 
     int val;​
     ListNode next;​
     ListNode() {}​
     ListNode(int val) { this.val = val; }​
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }​
}​



func Sort() {
	
}

func main() {

}



