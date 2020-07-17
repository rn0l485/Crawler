package worker

import (
	"errors"
	"net/http"
	"net/url"
)

type cookieJar struct {
	Jar					map[string][]*Cookie
}
func (j *cookieJar) SetCookies (u *url.URL, cookies []*http.Cookie){
	url := u.String()
	if given, ok := j.Jar[url]; ok {
		j.Jar[url] = append( given, cookies)
	} else {
		j.Jar[url] = cookies
	}
}

func (j *cookieJar) Cookies(u *url.URL) []*http.Cookie {
	url := u.String() 
	
	if given, ok := j.Jar[url]; ok {
		return given
	} else {
		return nil
	}
}