package routers

import (
	"net/http"
	"pkg/controllers"
)

//用户注册登录的路由
func registerUserRouter(mux *http.ServeMux) {
	//路由注册
	mux.HandleFunc("/login", controllers.Logig)
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/logout", controllers.Logout)
}
