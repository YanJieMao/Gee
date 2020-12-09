package gee

import (
	"net/http"
)

type H map[string]interface{}

type Contex struct {
	Writer http.ResponseWriter //远程对象
	Req    *http.Request

	Path   string //请求信息
	Method string

	StatusCode int //返回信息
}

func newContex(w http.ResponseWriter, req *http.Request) *Contex {
	return &Contex{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Contex) PostForm(key string) string {

	return c.Req.FormValue(key)
}

func (c *Contex) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Contex) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)

}

func (c *Contex) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)

}

func (c *Contex) String(code int, format string, values ...interface{}) {
	c.SetHeader("Context-type", "text/plain")

}
