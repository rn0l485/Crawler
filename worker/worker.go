package worker 

import (
	"bytes"
	"net/http"
	"io/ioutil"

	"github.com/json-iterator/go"
)


type Response struct {
	Status			string
	StatusCode 		int
	Proto			string
	Header 			http.Header
	Body 			[]byte	
}



type worker struct {
	Client 			*http.Client
	Header 			http.Header
	Cookies 		map[string][]*http.Cookie
}

func (w worker) Get(gurl string, channel chan Response) {

	// make req
	req, _ := http.NewRequest("GET", gurl, nil)
	if w.Header != nil {
		req.Header = w.Header
	}
	if cookie, ok := w.Cookies[gurl]; ok{
		for _,c := range cookie {
			req.AddCookie(c)
		}
	}

	// do request
	resp, err := w.Client.Do(req)
	if err != nil {
		panic(err)
	}

	// read body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// cookie
	cookies := resp.Cookies()
	w.SetCookies(gurl, cookies)

	// set resp
	given := Response{
		Status 		: resp.Status,
		StatusCode 	: resp.StatusCode,
		Proto		: resp.Proto,
		Header 		: resp.Header,
		Body 		: body,
	}

	channel <- given
}


func (w worker) Post(gurl string, data map[string]interface{}, channel chan Response) {
	// make req
	b, err := jsoniter.Marshal(data)
	if err != nil{
		panic(err)
	}
	req, _ := http.NewRequest( "POST", gurl, bytes.NewBuffer(b))
	if w.Header != nil {
		req.Header = w.Header
	}
	if cookie, ok := w.Cookies[gurl]; ok{
		for _,c := range cookie {
			req.AddCookie(c)
		}
	}

	// do req
	resp, err := w.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// cookie
	cookies := resp.Cookies()
	w.SetCookies(gurl, cookies)	

	given := Response{
		Status 		: resp.Status,
		StatusCode 	: resp.StatusCode,
		Proto		: resp.Proto,
		Header 		: resp.Header,
		Body 		: body,
	}

	channel <- given

}

func (w worker) SetCookies (url string, cookies []*http.Cookie){
	if given, ok := w.Cookies[url]; ok {
		for _,c := range cookies {
			w.Cookies[url] = append( given, c)
		}
	} else {
		w.Cookies[url] = cookies
	}
}


func InitWorker() worker{
	c := &http.Client{}
	w := worker{
		Client:c,
		Cookies: make(map[string][]*http.Cookie),
	}
	return w
}



















