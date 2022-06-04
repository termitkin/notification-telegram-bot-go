## notification-telegram-bot-go

The program accepts the text in the body and sends it to the telegram chat

[Docker image](https://hub.docker.com/r/termitkin/notification-telegram-bot-go)

### Example of docker-compose.yml

```yml
version: "3.8"

services:
  notification-telegram-bot-go:
    image: termitkin/notification-telegram-bot-go
    container_name: notification-telegram-bot-go
    restart: unless-stopped
    ports:
      - "5533:8080"
    environment:
      - TELEGRAM_BOT_TOKEN
      - TELEGRAM_BOT_CHAT_ID
    env_file:
      - .env
```

### Example of .env

```dotenv
TELEGRAM_BOT_TOKEN=token
TELEGRAM_BOT_CHAT_ID=chatId
```

### Example of nginx block

```nginx
# Notification to telegram
location ^~ /notification-to-telegram/bot-token {
    proxy_pass http://0.0.0.0:5533/;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection 'upgrade';
    proxy_set_header Host $host;
    proxy_cache_bypass $http_upgrade;
}
```


### Docker build

```bash
docker build -t termitkin/notification-telegram-bot-go:latest .
```

### Run docker container

```bash
docker pull termitkin/notification-telegram-bot-go:latest && docker-compose up -d
```
