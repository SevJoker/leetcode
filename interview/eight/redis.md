# redis


## 基础信息

#### redis基本数据类型

##### string对象
- raw 简单动态字符串
	- 大于 39 字节的字符串
- int long型数字
	- 8 个字节的长整型
- embstr embstr编码字符串
	- 小于等于 39 字节的字符串

##### list对象  
- ziplist压缩列表
	- 元素个数小于 512，且所有值都小于 64 字节
- linkedlist双端列表

##### 哈希对象 Hash
- ziplist压缩列表
	- 元素个数小于 512，且所有值都小于 64 字节
- hashtable字典
	- else


##### 集合对象 Set
- hashtable字典
	- else
- intset整数集合
	- 元素个数小于 512，且所有值都是整数


##### 有序集合对象
- ziplist   压缩字典
	- 元素个数小于 128，且所有值都小于 64 字节
- skiplist 跳跃表和字典
	- else