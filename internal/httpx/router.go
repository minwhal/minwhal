package httpx

import "net/http"

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{ServeMux: http.NewServeMux()}
}

func (r *Router) Handle(pattern string, handler AppHandler) {
	r.ServeMux.Handle(pattern, handler)
}
