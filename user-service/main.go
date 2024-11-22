package main

import (
	"net/http"

	. "github.com/Remoulding/12306-go/user-service/biz/userLogin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userLoginService := &UserLoginService{}

	router.POST("/api/user-service/v1/login", func(c *gin.Context) {
		var req UserLoginReqDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, Result[any]{Code: 400, Message: "Invalid request"})
			return
		}
		resp := userLoginService.Login(req)
		c.JSON(http.StatusOK, Success(&resp))
	})

	router.GET("/api/user-service/check-login", func(c *gin.Context) {
		accessToken := c.Query("accessToken")
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, Result[any]{Code: 400, Message: "Missing accessToken"})
			return
		}
		resp := userLoginService.CheckLogin(accessToken)
		c.JSON(http.StatusOK, Success(&resp))
	})

	router.GET("/api/user-service/logout", func(c *gin.Context) {
		accessToken := c.Query("accessToken")
		userLoginService.Logout(accessToken)
		c.JSON(http.StatusOK, Success[any](nil))
	})

	router.Run(":8080") // 启动服务，监听 8080 端口
}
