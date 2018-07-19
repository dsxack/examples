package main

import "github.com/valyala/fasthttp"

func httpHandlerHealth(_ *HttpServer, ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("health: ok")
}
