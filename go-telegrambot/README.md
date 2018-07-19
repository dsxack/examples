# golang telegram smotrovbot

HTTP шлюз для работы с телеграммом

### /message

Метод позволяет отправлять сообщения в чат по идетнификатору чата

Пример запроса:

```bash
curl -i -XPOST \
  "https://dev.dsxack.com/bots-telegram-go-smotrovbot/message" \
  -d '{"target_chat_id": -1001088291385, "message_text": "test message"}' \
```
