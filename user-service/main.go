package main

import "github.com/Remoulding/12306-go/user-service/configs"

func main() {
	configs.InitDBInstance()
	configs.InitCache()
	InitRpcServer()
	InitWebServer()
}
