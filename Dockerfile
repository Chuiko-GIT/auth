FROM golang:1.21-alpine AS builder

COPY . /github.com/Chuiko-GIT/chat/source/
WORKDIR /github.com/Chuiko-GIT/chat/source/

RUN go mod download
RUN go build -o ./bin/auth cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Chuiko-GIT/chat/source/bin/auth .

CMD ["./auth"]
