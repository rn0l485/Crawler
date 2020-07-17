package main

import (
	"fmt"
	"Crawler/worker"
)

func main() {
	w := worker.InitWorker()

	c := make(chan worker.Response)

	c <- w.Get("https://www.etron-valve.com")
	resp <- c

	fmt.Println(string(resp.Header))
}
