import Crawler
from collections import deque
from queue import Queue



class InitController():
	def __init__(self):
		self.__TaskOrders = deque([])
		self.__workers = {}

	def InitWorker(self, num=1):
		if type(num) != int :
			return "Please input int"
		for i in range(num):
			self.__workers[str(i)] = Crawler.InitCrawler()

	def SetTask(self, Tasks:list):
		for i in range(Tasks):
			if type(i) is Crawler.Task:
				self.__TaskOrders.appendleft(i)
			else:
				print("Error: Tasks[%s] type error!!!"%str(i))
				continue

	def Run(self):
		while True:
			pass
