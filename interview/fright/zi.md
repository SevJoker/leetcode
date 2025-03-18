linux
	线程通信
	进程通信
	socket与网络socket区别
	linux 获取锁的方式

设计
	短连接服务如何处理

中间件
	kafka消息不丢失的处理

	服务通信保证不丢失


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






mysql热点数据写入？
redis热点数据写入？

账户热点数据写入？


多机房之前的断联数据丢失问题 如何处理？



求123的全排列字符串






func A() {
	
}



<!--  自调用 -->
func B(a) []string{
	ret = []string{}
	if len(a) == 1 {
		return []sting{a}
	}

	tmp := B(a[1:])

	for item := range tmp {
		for i:=0;i<len(item);i++{
			ret = append(ret, item[0:i] + a[0] + item[i:])
		}
	}
	return ret
}




