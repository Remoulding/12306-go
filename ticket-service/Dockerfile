FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o service .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/service .
# COPY --from=builder /app/configs ./configs

# 暴露HTTP和gRPC端口
EXPOSE ${HTTP_PORT} ${GRPC_PORT}

CMD ["./service"]