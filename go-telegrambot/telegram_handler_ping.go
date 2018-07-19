package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"net"
	"time"
	"fmt"
	"strings"
	"github.com/tatsushid/go-fastping"
)

func telegramHandlePingMessage(service *TelegramBot, message *tgbotapi.Message) {
	var err error

	defer func() {
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
				"Error to ping site: %v", err,
			))
			service.Send(msg)
		}
	}()

	fields := strings.Fields(message.Text)
	site := fields[1]
	duration := time.Duration(0)

	pinger := fastping.NewPinger()
	addr, err := net.ResolveIPAddr("ip4:icmp", site)
	if err != nil {
		return
	}

	recv := make(chan time.Duration)
	pinger.AddIPAddr(addr)
	pinger.OnRecv = func(addr *net.IPAddr, duration time.Duration) {
		recv <- duration
	}
	go pinger.Run()
	duration = <-recv

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"%s ponged with %s", addr.String(), duration,
	))

	service.Send(msg)

	return
}
