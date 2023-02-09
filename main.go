package main

import (
	"log"
	"net/http"
	"pkg/routers"
)

func main() {
	mux := http.NewServeMux()
	//ServeMux是一个结构体
	//参数2 是个接口  实现了ServeHTTP(ResponseWriter, *Request)

	//启动路由
	routers.RegisterSocialRouter(mux)
	log.Fatal(http.ListenAndServe(":8888", mux))
}
