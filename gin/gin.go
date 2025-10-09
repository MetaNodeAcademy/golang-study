package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello gin",
			"code":    200,
			"data": map[string]interface{}{
				"name": "gin",
				"age":  18,
			},
		})
	})

	//普通路由
	router.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"code":    200,
		})
	})

	//分组路由(路由组)
	v1 := router.Group("v1") //http://localhost:8888/v1/user
	v1.GET("/user", func(con *gin.Context) {
		con.JSON(http.StatusOK, gin.H{
			"message": "用户列表",
			"code":    200,
		})
	})
	v1.POST("/userInfo", func(con *gin.Context) {
		//JSON输出
		// con.JSON(http.StatusOK, gin.H{
		// 	"message": "用户信息",
		// 	"code":    200,
		// })

		//xml输出
		con.XML(http.StatusOK, gin.H{
			"message": "用户信息",
			"code":    200,
		})

	})

	v1.DELETE("/deleteUser", deleteUser)

	//重定向到外部链接
	router.GET("/redi", func(con *gin.Context) {
		con.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	//重定向到内部链接
	router.GET("/redi2", func(con *gin.Context) {
		con.Redirect(http.StatusFound, "/user")
	})

	//读取静态文件
	router.Static("/static", "./static") //目录
	router.StaticFile("/aaa.txt", "./static/aaa.txt")

	router.Run(":8888")
}

func deleteUser(con *gin.Context) {
	fmt.Println("删除用户")
}
