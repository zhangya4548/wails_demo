package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// App struct
type App struct {
	ctx context.Context
	h *http.Server
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

type Home struct {
	Uri string `json:"uri"`
	Dir string `json:"dir"`
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx

	// 实例化gin
	r := gin.New()
	//开启中间件 允许使用跨域请求
	r.Use(Cors())

	//模型绑定/参数验证 curl -g "http://localhost:9999/login" -H "content-type: application/json" -X POST -d '{ "user": "admin", "password":"123456" }'
	r.POST("/home", func(c *gin.Context) {
		var home Home
		if err := c.ShouldBindJSON(&home); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		// if home.Uri != "admin" || home.Dir != "123456"{
		// 	c.JSON(http.StatusBadRequest, gin.H{"error":"参数错误"})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{"msg":"成功","data":home.Uri})
	})

	srv := &http.Server{
		Addr:    ":9999",
		Handler: r,
	}
	a.h = srv

	go func() {
		if err := a.h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("gin监听: %s\n", err)
		}
	}()
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
	if err := a.h.Shutdown(ctx); err != nil {
		log.Fatal("gin退出异常:", err)
	}
}



// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}