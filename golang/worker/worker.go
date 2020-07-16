package worker 

import (
	"log"
	"net/http"
	"io/ioutil"
)


type response struct {
	Status			string
	StatusCode 		int
	Proto			string
	Header 			*http.Header
	Body 			string	
}



type worker struct {
	client 			*http.Client
}

func (w worker) Get(url string) *response {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := w.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	given := response{
		Status 		: resp.Status,
		StatusCode 	: resp.StatusCode,
		Proto		: resp.Proto,
		Header 		: resp.Header,
		Body 		: body,
	}

	return given
}

func (w worker) Post(url string, data map[string]interface{}) {
	req, _ := http.NewRequestWithContext(url, )
}






















