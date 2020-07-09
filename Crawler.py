import requests


class InitCrawler():
	def __init__(self):
		self.__Session = requests.Session() 

	def Get(self, Task ):
		r = self.__Session.get(Task.url, headers=Task.Header)
		if r.status_code == requests.codes.ok:
			Task.resp = r.text 
		else:
			Task.resp = r.status_code
		return Task


	def Post(self, Task ):
		r = self.__Session.post(Task.url, headers = Task.Header, data = Task.dataset)
		if r.status_code == requests.codes.ok:
			Task.resp = r.text 
		else:
			Task.resp = r.status_code
		return Task


class Task():
	def __init__(self, Header, url:str, dataset=None):
		self.Header = Header
		self.url = url
		self.dataset = dataset
		self.resp = None

