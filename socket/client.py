#! /usr/bin/env python
# -*- coding: utf-8 -*-
#load-yaml.py 
#客户端
import socket
import os
import time
import threading
client=socket.socket()
client.connect(('127.0.0.1', 8083))

pid = os.getpid()
while 1:
	res = raw_input("client:")
	# res=('%s hello' %os.getpid()).encode('utf-8')
	client.send( (str(pid)+"_client:"+str(res)).encode('utf-8'))
	data=client.recv(1024)
	print(data.decode('utf-8'))