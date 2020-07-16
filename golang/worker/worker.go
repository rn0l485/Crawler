package worker 

import (
	"log"
	"bytes"
	"net/http"
	"io/ioutil"

	"github.com/json-iterator/go"
)


type Response struct {
	Status			string
	StatusCode 		int
	Proto			string
	Header 			*http.Header
	Body 			[]byte	
}



type Worker struct {
	client 			*http.Client
	Header 			*http.Header
}

func (w worker) Get(url string) *response {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = w.Header
	resp, err := w.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	given := Response{
		Status 		: resp.Status,
		StatusCode 	: resp.StatusCode,
		Proto		: resp.Proto,
		Header 		: resp.Header,
		Body 		: body,
	}

	return given
}

func (w worker) Post(url string, data map[string]interface{}) {
	b, err := jsoniter.Marshal(data)
	if err != nil{
		panic(err)
	}

	req, _ := http.NewRequest( "POST", url, bytes.NewBuffer(b))
	req.Header = w.Header
	resp, err := w.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	given := Response{
		Status 		: resp.Status,
		StatusCode 	: resp.StatusCode,
		Proto		: resp.Proto,
		Header 		: resp.Header,
		Body 		: body,
	}

	return given

}


func InitWorker() (w *worker){
	c := &http.Client{}
	w := Worker{
		client:c,
	}
	return 
}



















