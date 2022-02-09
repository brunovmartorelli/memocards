package router

import (
	"github.com/valyala/fasthttp"
)

func ParsePaths(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		//ctx.Response.Header.Set("Access-Control-Allow-Credentials", "*")
		//ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

		next(ctx)
	}
}
