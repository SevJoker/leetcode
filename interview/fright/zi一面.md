ysmysql acid是啥？ 各自实现是怎么做到哦
map 数据结构是啥？  chan的数据结构是啥

redis 备份原理是啥  aof 文件会越来越大  redis 怎么优化的

golang的相关数据结构实现 比如 map等

zset 跳表结构 详细说说

标题
删除链表重复元素

题目描述
存在一个按升序排列的单向链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中【没有重复出现】的数字，返回同样按升序排列的结果链表。​

输入  1->1->1->2->3->4->4​
输出   2->3​

public class ListNode {​
     int val;​
     ListNode next;​
     ListNode() {}​
     ListNode(int val) { this.val = val; }​
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }​
}​



