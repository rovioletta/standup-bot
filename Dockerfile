FROM golang:1.27-alpine as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o bin/standup-bot ./cmd/standup-bot

FROM alpine

COPY --from=builder /app/bin/standup-bot /standup-bot
COPY --from=builder /app/.env .env

CMD [ "standup-bot" ]