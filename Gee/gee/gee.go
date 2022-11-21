package gee

import (
	"net/http"
)

// 定义一个普遍使用的函数类型，避免后面再次定义
type HandlerFunc func(*Context)

// 定义路由表
type Engine struct {
	router *router
}

// 工厂模式的构造方法，返回一个实例
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// 实现GET方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

// 实现POST方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

// 实现Run方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 完成统一的控制入口方法ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
