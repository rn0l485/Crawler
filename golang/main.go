package main

import (
	"fmt"
	"golang/worker"
)

func main() {
	w := worker.InitWorker()
	resp := w.Get("https://www.etron-valve.com")
	fmt.Println(string(resp.Body))
}