package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func startGinServer() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 定义一个简单的 HTTP 路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin!",
		})
	})

	// 启动 Gin HTTP 服务
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start Gin server:", err)
	}
}
