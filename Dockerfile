FROM golang:1.18.3-alpine as build

WORKDIR /app

COPY app ./app
COPY go.mod ./

WORKDIR /app/app

RUN go build -o /notification-telegram-bot-go

FROM alpine:latest 

COPY --from=build /notification-telegram-bot-go ./

CMD [ "/notification-telegram-bot-go" ]
