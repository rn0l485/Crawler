package worker

import (
	"net/http"
	"net/url"
)

type cookieJar struct {

}
func (j *cookieJar) SetCookies (u *url.URL, cookies []*http.Cookie){
	url := u.String()
	if given, ok := j.Jar[url]; ok {
		for _,c := range cookies {
			j.Jar[url] = append( given, c)
		}
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
