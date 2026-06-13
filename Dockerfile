FROM golang:1.25-alpine AS builder
LABEL authors="aaron"
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o api ./cmd/main.go
RUN go build -o worker ./cmd/worker/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/api .
COPY --from=builder /app/worker .

COPY .env .

EXPOSE 8080

# Default to API
CMD ["./api"]
