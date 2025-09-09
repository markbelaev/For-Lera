FROM golang:1.25.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lera_bot ./cmd/bot

FROM alpine:3.22.1
WORKDIR /app
COPY --from=builder /app/lera_bot .
CMD ["./lera_bot"]
