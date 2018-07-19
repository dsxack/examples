package main

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
)

func httpHandlerMessage(server *HttpServer, ctx *fasthttp.RequestCtx) {
	var err error

	defer func() {
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusServiceUnavailable)
		}
	}()

	request := &struct {
		TargetChatID int64  `json:"target_chat_id"`
		MessageText  string `json:"message_text"`
	}{}

	err = json.Unmarshal(ctx.Request.Body(), request)
	if err != nil {
		return
	}

	var msg tgbotapi.Message
	msg, err = server.telegramBot.Send(
		tgbotapi.NewMessage(request.TargetChatID, request.MessageText),
	)
	if err != nil {
		return
	}

	var msgBytes []byte
	msgBytes, err = json.Marshal(msg)
	if err != nil {
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(msgBytes)
}
