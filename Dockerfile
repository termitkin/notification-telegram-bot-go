FROM golang:1.18-alpine

WORKDIR /app

COPY ./app/*.go ./
COPY ./go.mod ./

RUN go build -o /notification-telegram-bot-go

CMD [ "/notification-telegram-bot-go" ]
