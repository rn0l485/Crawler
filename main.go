package main

import (
	"fmt"
	"Crawler/worker"
)

func main() {
	w := worker.InitWorker()

	c := make(chan worker.Response)

	go w.Get("https://www.etron-valve.com", c)
	go w.Get("https://michaelchen.tech/golang-programming/concurrency/", c)
	for i:=0; i<2; i++ {
		resp := <- c
		fmt.Println(resp.Header)
	}
}
