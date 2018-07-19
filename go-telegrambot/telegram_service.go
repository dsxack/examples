package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

type TelegramBot struct {
	*tgbotapi.BotAPI
}

func (bot *TelegramBot) ListenUpdates() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		go bot.handleUpdate(update)
	}

	return nil
}

func (bot *TelegramBot) handleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		log.Printf("%v++", update.Message)

		fields := strings.Fields(update.Message.Text)
		switch fields[0] {
		case "/ping":
			telegramHandlePingMessage(bot, update.Message)
		}

	}
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	botapi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	botapi.Debug = true

	log.Printf("Authorized on account %s", botapi.Self.UserName)

	return &TelegramBot{BotAPI: botapi}, nil
}
