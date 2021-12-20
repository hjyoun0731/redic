package redic

// #cgo LDFLAGS: -L ./lib -lhiredis
//
// #include "./lib/redic.c"
import "C"

import (
	"errors"
	"strings"
	"unsafe"
)

type Client struct {
	Conn *C.redisContext
	host string
	port int
}

func NewClient(host string, port int) (*Client, error) {
	t := &Client{host: host, port: port}
	c := C.NewConnect(C.CString(t.host), C.int(t.port))
	if c == nil {
		return nil, errors.New("Connection error")
	}
	t.Conn = c
	return t, nil
}

func (t *Client) Close() {
	C.CloseConnect(t.Conn)
}

func (t *Client) RedisCommand(arg ...string) string {
	strings.Join(arg, " ")

	p := C.RedisCommand(t.Conn, C.CString(strings.Join(arg, " ")))
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))

	return s
}
