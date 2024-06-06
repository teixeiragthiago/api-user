FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-user ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api-user .

COPY cmd/.env .env

EXPOSE 5000

CMD ["./api-user"]
