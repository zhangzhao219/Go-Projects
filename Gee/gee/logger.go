package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()                                                            // 开始计时
		c.Next()                                                                   // 等待用户自己的Handler处理结束
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t)) // 打印时间
	}
}
