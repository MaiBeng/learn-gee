package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// 存放路由与方法的映射
type Engine struct {
	route map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{route : make(map[string]HandlerFunc)}
}

// 注册get方法
func (engine *Engine) Get(url string, handler HandlerFunc) {
	engine.route[url + "-GET"] = handler
}

// 注册post方法
func (engine *Engine) Post(url string, handler HandlerFunc) {
	engine.route[url + "-POST"] = handler
}

// 启动并运行服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// Engine实例实现了ServeHTTP方法，作为ListenAndServe的第二参数
func (engine *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	key := request.URL.Path + "-" + request.Method

	if router, ok := engine.route[key]; ok {
		router(response, request)
	}else {
		fmt.Println("404")
	}
}