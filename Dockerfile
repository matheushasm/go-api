FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app

COPY .env.example .env

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
