FROM golang:1.23.3

WORKDIR /app

# 复制服务代码和本地模块
COPY . .
COPY ../idl-gen ./idl-gen

# 设置 Go 环境变量
ENV GO111MODULE=on

WORKDIR /app/user-service

RUN go mod tidy
RUN go build -o user-service

CMD ["./user-service"]
