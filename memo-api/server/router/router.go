package router

import (
	fhr "github.com/fasthttp/router"
)

type Router struct {
	Router *fhr.Router
}

func New() *Router {
	r := &Router{
		Router: fhr.New(),
	}
	return r
}
