package worker 

import (
	"regexp"
)

type Reader interface {
	Read(resp *Response) ( targetItem []byte, nextUrl string, err error)
}

