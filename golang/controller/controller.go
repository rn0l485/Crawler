package controller


import (
	"github.com/rn0l485/Crawler/worker"
)

type controller struct {
	multiChannel 		[]chan worker.Response
	multiWorker 		[]worker.Worker
}

func (c *controller) SetWorker(n int) {
	
}