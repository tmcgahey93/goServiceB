FROM golang:1.23.5 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goserviceb .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/goserviceb .

EXPOSE 8081

RUN chmod +x goserviceb

CMD ["./goserviceb"]