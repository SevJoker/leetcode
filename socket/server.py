#! /usr/bin/env python
# -*- coding: utf-8 -*-
#load-yaml.py 

import socket
import time
from multiprocessing import Pool,Lock,Process


from socket import *
import select


def init():
	pass


def selectServer():
	server = socket(AF_INET, SOCK_STREAM)
	server.bind(('127.0.0.1',8083))
	server.listen(5)
	server.setblocking(False)
	# 初始化将服务端socket对象加入监听列表，后面还要动态添加一些conn连接对象，当accept的时候sk就有感应，当recv的时候conn就有动静
	rlist=[server,]
	rdata = {}  #存放客户端发送过来的消息

	wlist=[]  #等待写对象
	wdata={}  #存放要返回给客户端的消息

	print '预备！监听！！！'
	count = 0 #写着计数用的，为了看实验效果用的，没用
	while True:
		# 开始 select 监听,对rlist中的服务端server进行监听，select函数阻塞进程，直到rlist中的套接字被触发（在此例中，套接字接收到客户端发来的握手信号，从而变得可读，满足select函数的“可读”条件），被触发的（有动静的）套接字（服务器套接字）返回给了rl这个返回值里面；
		rl,wl,xl=select.select(rlist,wlist,[],0.5)
		print '次数>>',(count),rl,wl,xl
		count = count + 1
		# 对rl进行循环判断是否有客户端连接进来,当有客户端连接进来时select将触发
		for sock in rl:
			# 判断当前触发的是不是socket对象, 当触发的对象是socket对象时,说明有新客户端accept连接进来了
			if sock == server:
				# 接收客户端的连接, 获取客户端对象和客户端地址信息
				conn,addr=sock.accept()
				#把新的客户端连接加入到监听列表中，当客户端的连接有接收消息的时候，select将被触发，会知道这个连接有动静，有消息，那么返回给rl这个返回值列表里面。
				rlist.append(conn)
			else:
				# 由于客户端连接进来时socket接收客户端连接请求，将客户端连接加入到了监听列表中(rlist)，客户端发送消息的时候这个连接将触发
				# 所以判断是否是客户端连接对象触发
				try:
					data=sock.recv(1024)
					#没有数据的时候，我们将这个连接关闭掉，并从监听列表中移除
					if not data:
						sock.close()
						rlist.remove(sock)
						continue
					print "received "+data.decode()+" from client "+ str(sock)
					#将接受到的客户端的消息保存下来
					rdata[sock] = data.decode()
					#将客户端连接对象和这个对象接收到的消息加工成返回消息，并添加到wdata这个字典里面
					wdata[sock]=data.upper()
					#需要给这个客户端回复消息的时候，我们将这个连接添加到wlist写监听列表中
					wlist.append(sock)
				#如果这个连接出错了，客户端暴力断开了（注意，我还没有接收他的消息，或者接收他的消息的过程中出错了）
				except Exception:
					#关闭这个连接
					sock.close()
					#在监听列表中将他移除，因为不管什么原因，它毕竟是断开了，没必要再监听它了
					rlist.remove(sock)
		# 如果现在没有客户端请求连接,也没有客户端发送消息时，开始对发送消息列表进行处理，是否需要发送消息
		for sock in wl:
			sock.send(wdata[sock])
			wlist.remove(sock)
			wdata.pop(sock)

		# #将一次select监听列表中有接收数据的conn对象所接收到的消息打印一下
		# for k,v in rdata.items():
		#	 print(k,'发来的消息是：',v)
		# #清空接收到的消息
		# rdata.clear()




def noBlockServer():
	server=socket.socket()
	server.setsockopt(socket.SOL_SOCKET,socket.SO_REUSEADDR,1)
	server.bind(('127.0.0.1',8083))
	server.listen(5)

	server.setblocking(False) #设置不阻塞
	r_list=[]  #用来存储所有来请求server端的conn连接
	w_list={}  #用来存储所有已经有了请求数据的conn的请求数据

	while 1:
		try:
			conn,addr=server.accept() #不阻塞，会报错
			r_list.append(conn)  #为了将连接保存起来，不然下次循环的时候，上一次的连接就没有了
		except Exception as e:
			# 强调强调强调：！！！非阻塞IO的精髓在于完全没有阻塞！！！
			# time.sleep(0.5) # 打开该行注释纯属为了方便查看效果
			print e
			print "在做其他的事情"
			print "rlist: ",len(r_list)
			print "wlist: ",len(w_list)
		# 遍历读列表，依次取出套接字读取内容
		del_rlist=[] #用来存储删除的conn连接
		for conn in r_list:
			try:
				data=conn.recv(1024) #不阻塞，会报错
				if not data: #当一个客户端暴力关闭的时候，会一直接收b''，别忘了判断一下数据
					conn.close()
					del_rlist.append(conn)
					continue
				w_list[conn]=data.upper()
			except Exception as e: # 没有收成功，则继续检索下一个套接字的接收
				continue

		# 遍历写列表，依次取出套接字发送内容
		del_wlist=[]
		for conn,data in w_list.items():
			try:
				conn.send(data)
				del_wlist.append(conn)
			except BlockingIOError:
				continue
		# 清理无用的套接字,无需再监听它们的IO操作
		for conn in del_rlist:
			r_list.remove(conn)
		#del_rlist.clear() #清空列表中保存的已经删除的内容
		for conn in del_wlist:
			w_list.pop(conn)
		#del_wlist.clear()

# 阻塞多进程模型
def blockServer():
	server = socket.socket()
	# server.setsockopt(socket.SOL_SOCKET,socket.SO_REUSEADDR,1)
	server.bind(('127.0.0.1',8083))
	server.listen(5)
	print "blockServer start"

	while 1:
		conn,addr=server.accept()
		print "server conn:"+ str(addr)

		p = Process(target=deal_conn, args=(conn,addr,))
		p.start()
		# r = pool.apply_async(deal_conn, args=(conn,addr,))
		# while 1:
		# 	data=conn.recv(1024)
		# 	print "server revc:"+data
		# 	conn.send((str(data)+str(addr)).encode("utf-8"))

def deal_conn(conn,addr):
	print "deal_conn:",str(conn),str(addr)
	while 1:
		data=conn.recv(1024)
		print "server revc:"+data
		conn.send((str(data)+str(addr)).encode("utf-8"))


if __name__ == '__main__':
	
	# blockServer()
	# noBlockServer()
	selectServer()
