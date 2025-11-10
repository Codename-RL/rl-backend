FROM golang:1.25.3-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine:latest
RUN apk add --no-cache tzdata

WORKDIR /root/
COPY --from=builder /app/app .
COPY config.json .
CMD ["./app"]
