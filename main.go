package main

import (
	"fmt"
	"github.com/mahongcheng/go-gin-study/pkg/setting"
	"github.com/mahongcheng/go-gin-study/routers"
	"net/http"
)

//编写项目启动、路由文件

func main() {
	//1.返回 Gin 的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
	//router := gin.Default()
	////2.创建不同的 HTTP 方法绑定到Handlers中，也支持 POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
	//router.GET("/test", func(c *gin.Context) {
	//	//3.gin.H{…}：就是一个map[string]interface{}
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	//4.gin.Context：Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，
	// 在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等

	//Addr：监听的 TCP 地址，格式为:8000
	//Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
	//TLSConfig：安全传输层协议（TLS）的配置
	//ReadTimeout：允许读取的最大时间
	//ReadHeaderTimeout：允许读取请求头的最大时间
	//WriteTimeout：允许写入的最大时间
	//IdleTimeout：等待的最大时间
	//MaxHeaderBytes：请求头的最大字节数
	//ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
	//ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出）

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//开始监听服务，监听 TCP 网络地址，Addr 和调用应用程序处理连接上的请求
	s.ListenAndServe()
}
