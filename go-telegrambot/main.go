package main

const (
    secret = "secret key"
	servicePort = 80
)

func main() {
	errChan := make(chan error)

	// telegram bot
	var telegramBot *TelegramBot
	{
		var err error

		telegramBot, err = NewTelegramBot(secret)
		if err != nil {
			panic(err)
		}

		go func() {
			errChan <- telegramBot.ListenUpdates()
		}()
	}

	// http rest
	{
		httpServer := NewHttpServer(telegramBot)
		go func() {
			errChan <- httpServer.Listen(servicePort)
		}()
	}

	select {
	case err := <-errChan:
		panic(err)
	}
}
