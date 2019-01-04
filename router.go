package udshttp

import (
	"net/http"
)

const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)

var RouterMethods = []string{
	MethodGet,
	MethodHead,
	MethodPost,
	MethodPut,
	MethodPatch,
	MethodDelete,
	MethodConnect,
	MethodOptions,
	MethodTrace,
}

type Handler func(resp http.ResponseWriter, req *http.Request)

type Router struct {
	//key1-method, key2-path
	routers map[string]map[string]Handler
}

func NewRouter() *Router {
	router := &Router{
		routers: make(map[string]map[string]Handler),
	}
	for _, method := range RouterMethods {
		router.routers[method] = make(map[string]Handler)
	}
	return router
}

func (r *Router) Get(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodGet][path] = handler
}

func (r *Router) Head(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodHead][path] = handler
}

func (r *Router) Post(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodPost][path] = handler
}

func (r *Router) Put(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodPut][path] = handler
}

func (r *Router) Patch(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodPatch][path] = handler
}

func (r *Router) Delete(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodDelete][path] = handler
}

func (r *Router) Connect(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodConnect][path] = handler
}

func (r *Router) Options(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodOptions][path] = handler
}

func (r *Router) Trace(path string, handler Handler) {
	if path == "" || handler == nil {
		return
	}
	//cover if exist
	r.routers[MethodTrace][path] = handler
}

func (r *Router) match(method, path string) Handler {
	if v1, ok := r.routers[method]; ok {
		if handle, ok := v1[path]; ok {
			return handle
		}
	}
	//if not match any handler, NotFound will be return
	return http.NotFound
}
