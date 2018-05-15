package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println("Handled message: " + string(msg))
		m.Broadcast(msg)
	})

	go func() {
		everySecond := time.Tick(time.Second)

		select {
		case <-everySecond:
			m.Broadcast([]byte(`{"message": "Every second tick"}`))
		}
	}()

	r.Run("0.0.0.0:5000")
}
