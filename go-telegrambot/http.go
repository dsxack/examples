package main

import (
	"github.com/valyala/fasthttp"
	"fmt"
)

type HttpServer struct {
	telegramBot *TelegramBot
}

func (server *HttpServer) Listen(port int) error {
	return fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/message":
			httpHandlerMessage(server, ctx)

		case "/health":
			httpHandlerHealth(server, ctx)
		}
	}))
}

func NewHttpServer(telegramBot *TelegramBot) *HttpServer {
	return &HttpServer{telegramBot}
}
