package worker 

import (
	"bytes"
	"net/http"
	"net/url"
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

func (w worker) Get(gurl string) *Response {
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

	return *given
}


func (w worker) Post(gurl string, data map[string]interface{}) *Response{
	// make req
	b, err := jsoniter.Marshal(data)
	if err != nil{
		panic(err)
	}
	req, _ := http.NewRequest( "POST", gurl, bytes.NewBuffer(b))
	if w.Header != nil {
		req.Header = w.Header
	}
	if given, ok := w.Client.Jar.Jar[gurl]; ok{
		for _,c := range w.Client.Jar.Jar[gurl] {
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
	w.SetCookies(u, cookies)	

	given := Response{
		Status 		: resp.Status,
		StatusCode 	: resp.StatusCode,
		Proto		: resp.Proto,
		Header 		: resp.Header,
		Body 		: body,
	}

	return *given

}

func (w worker) SetCookies (u string, cookies []*http.Cookie){
	if given, ok := w.Cookies[url]; ok {
		for _,c := range cookies {
			w.Cookies[url] = append( given, c)
		}
	} else {
		w.Cookies[url] = cookies
	}
}

func (w worker) Cookies(u string) []*http.Cookie {	
	if given, ok := w.Cookies[url]; ok {
		return given
	} else {
		return nil
	}
}

func InitWorker() *worker{
	c := &http.Client{}
	return worker{Client:c,}
}



















