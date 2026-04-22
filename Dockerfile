FROM golang:1.25-alpine AS builder
LABEL authors="aaron"
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

EXPOSE 8080

CMD ["./main"]