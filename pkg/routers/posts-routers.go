package routers

import (
	"net/http"
	"pkg/controllers"
)

func registerPostRouter(mux *http.ServeMux) {
	mux.HandleFunc("/create", controllers.Create)
}
