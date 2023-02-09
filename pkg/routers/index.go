package routers

import (
	"net/http"
)

//用户总路由
var RegisterSocialRouter = func(mux *http.ServeMux) {
	//users路由
	registerUserRouter(mux)
	//Posts路由
	registerPostRouter(mux)
}
