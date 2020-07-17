package main

import (
	"fmt"
	"Crawler/worker"
)

func main() {
	w := worker.InitWorker()

	c := make(chan worker.Response)

	go w.Get("https://www.etron-valve.com", c)

	resp := <- c

	fmt.Println(resp.Header)
}
