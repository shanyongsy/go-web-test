package gin

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	gorm_mgr "github.com/shanyongsy/go-web-test/gorm"
)

type MutexWrapper struct {
	mu sync.Mutex
}

var (
	logFileName string
	lock        sync.Mutex
)

// 启动 Gin 框架 Web 服务
func StartGinServer(port string, https bool, db *gorm_mgr.DBManager) (*http.Server, error) {
	gin.DisableConsoleColor()
	logName := "./log/gin.log"
	f, _ := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(CustomLogger())
	r.Use(gin.Recovery())

	// 使用中间件将 DBManager 传递给每个路由
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// r.POST("/recharge", handleInterRecharge)
	// r.POST("/simple-recharge", handleInterSimpleRecharge)
	// r.POST("/check", handleInterCheck)
	// r.POST("/status-change", handelChangeStatus)
	r.POST("/vc/notice", handleInterRechargeFromSpider)
	r.GET("/ping", handlePing)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if https {
			if err := server.ListenAndServeTLS("./openssl/server.crt", "./openssl/server.key"); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Error starting server: %v", err)
			}
		} else {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Error starting server: %v", err)
			}
		}
	}()

	return server, nil
}

// CustomLogger 自定义日志中间件
func CustomLogger() gin.HandlerFunc {
	// 打开或创建日志文件
	file, err := os.OpenFile("./log/"+generateLogFileName(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
	}

	// 将标准库的 log 输出到文件
	log.SetOutput(file)

	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求

		c.Next()

		// 记录自定义的日志信息到日志
		log.Printf("Custom log: method=%s path=%s ip=%s status=%d latency=%s\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

// generateLogFileName 生成日志文件名
func generateLogFileName() string {
	// return "app_" + time.Now().Format("20060102150405") + ".log"
	return "app_" + time.Now().Format("20060102") + ".log"
}
