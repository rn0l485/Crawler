# Crawler
Crawler in Golang


Easy to use in goroutine as below:
```golang
func main(){
	w := worker.InitWorker()

	c := make(chan worker.Response)

	go w.Get("First web url", c)
	go w.Get("Second web url", c)
	for i:=0; i<2; i++ {
		resp := <- c
		fmt.Println(resp.Header)
	}
}
```

The Response is constructed as below:

```golang 
type Response struct {
	Status			string
	StatusCode 		int
	Proto			string
	Header 			http.Header
	Body 			[]byte	
}
```


Currently, just support GET and POST, the useful methods I guess. 

For the worker struct:
```golang
type worker struct {
	Client 			*http.Client
	Header 			http.Header
	Cookies 		map[string][]*http.Cookie
}
```

For the GET method:
```golang
func (w worker) Get(url string) Response
```
For the POST method:
```golang
func (w worker) Post(url string, data map[string]interface{}) Response
```

