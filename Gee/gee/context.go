package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 给map[string]interface{}起了一个别名gee.H，构建JSON数据时，显得更简洁。
type H map[string]interface{}

type Context struct {
	// 原始的两个参数
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求信息
	Path   string
	Method string
	Params map[string]string
	// 响应信息
	StatusCode int
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 创建一个Context实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 根据key返回用户输入的value,属于POST方法的工具
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 根据key返回用户输入的value,属于GET方法的工具
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 写入状态码并更改Context的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 帮助下面的方法快速构造响应
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// 构造字符串类型的响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 构造JSON类型的响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer) // 流数据构造json
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 构造data类型的响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 构造HTML类型的响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
