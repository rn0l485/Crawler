package worker 





type Reader interface {
	Read(resp *Response, targetTag string) ( targetItem []byte, nextUrl string, err error)
}

