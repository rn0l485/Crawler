# Crawler
Crawler in Golang


Easy to use in goroutine as below:
```golang
	w := worker.InitWorker()

	c := make(chan worker.Response)

	c <- w.Get("your url here")
	resp <- c

	fmt.Println(resp.Header)
```

Currently, just support GET and POST, the useful methods I guess. 

For the GET method:

func (w worker) Get(url string) Response

For the POST method:

func (w worker) Post(url string, data map[string]interface{}) Response


