package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	function "github.com/shanyongsy/go-web-test/func"
	gin_mgr "github.com/shanyongsy/go-web-test/gin"
)

func main() {
	function.PreStart()

	https := flag.Bool("https", false, "use HTTPS")
	port := flag.String("port", "8080", "server port")

	// dbHost := flag.String("dbhost", "localhost", "database host")
	// dbPort := flag.String("dbport", "3306", "database port")
	// dbUser := flag.String("dbuser", "root", "database user")
	// dbPwd := flag.String("dbpwd", "123456", "database password")
	// dbName := flag.String("dbname", "recharge_db", "database name")

	flag.Parse()
	fmt.Printf("is https:%v,\t listen port:%v\n", *https, *port)
	// fmt.Printf("mysql host:%v,\t port:%v,\t user:%v,\t pwd:%v\n", *dbHost, *dbPort, *dbUser, *dbPwd)
	fmt.Println("server is starting...")

	// 初始化数据库
	// dbManager, err := gorm_mgr.NewDBManager(*dbHost, *dbPort, *dbUser, *dbPwd, *dbName)
	// if err != nil {
	// 	log.Fatalf("Error connecting to database: %v", err)
	// }

	// 初始化数据库
	// dbManager, err := dbopt.NewDBManager(*dbHost, *dbPort, *dbUser, *dbPwd, *dbName)

	// 启动 Gin 框架 Web 服务
	httpServer, err := gin_mgr.StartGinServer(*port, *https, nil)
	if err != nil {
		log.Fatalf("Error starting Gin server: %v", err)
	} else {
		fmt.Println("server is running...")
	}

	// 使用 WaitGroup 等待所有后台线程完成
	var wg sync.WaitGroup

	// 退出信号通道
	exitChan := make(chan struct{})

	// 启动后台处理线程1
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	backthread.ProcessRecharge(dbManager, exitChan)
	// }()

	// 启动后台处理线程2
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	backthread.ProcessRecharge(exitChan)
	// }()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGUSR1, syscall.SIGUSR2)

	select {
	case sig := <-sigChan:
		fmt.Println("sig:", sig)
		fmt.Println("Received exit signal. Performing cleanup...")

		// 发送退出信号给后台线程
		close(exitChan)

		// 在这里执行清理操作，例如关闭数据库连接等
		httpServer.Close()
		fmt.Println("http server closed")

		// 等待所有后台线程完成
		wg.Wait()

		// 关闭数据库连接
		// dbManager.Close()
		// fmt.Println("db manager closed")

		fmt.Println("Cleanup complete. Exiting...")
	}

	fmt.Println("server exited")

}
