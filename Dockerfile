FROM golang:1.18-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./*.go ./

RUN go build -o /notification-telegram-bot-go

CMD [ "/notification-telegram-bot-go" ]
