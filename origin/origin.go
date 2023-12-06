package origin

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/google/uuid"
)

// Data struct to represent the JSON data
type Data struct {
	ServerID  string `json:"serverid"`
	ValueID   string `json:"valueid"`
	AccountID string `json:"accountid"`
	OrderID   string `json:"orderid"`
}

type MutexWrapper struct {
	mu sync.Mutex
}

// ServerConfig 包含HTTP服务器的配置
type ServerConfig struct {
	Address   string     // 服务器地址
	TLSConfig *TLSConfig // TLS配置
}

// TLSConfig 包含服务器端证书验证所需的配置
type TLSConfig struct {
	CACertFile     string // CA证书文件路径
	ServerCertFile string // 服务器证书文件路径
	ServerKeyFile  string // 服务器私钥文件路径
}

func main() {
	https := flag.Bool("https", false, "use HTTPS")
	flag.Parse()
	fmt.Println("https:", *https)

	fmt.Println("server is starting...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGUSR1, syscall.SIGUSR2)

	dataStore := make(map[string]Data)
	mu := &MutexWrapper{} // 创建一个互斥锁

	if !*https {
		go func() {
			// HTTP监听
			httpServerAddr := ":8080"
			httpConfig := &ServerConfig{Address: httpServerAddr}
			StartServer(httpConfig, dataStore, mu)
		}()
	} else {
		// go func() {
		// 	// HTTPS监听（CA签名证书）
		// 	httpServerAddr := ":8081"
		// 	httpsCACertConfig := &ServerConfig{
		// 		Address: httpServerAddr,
		// 		TLSConfig: &TLSConfig{
		// 			CACertFile:     "./ca.crt",        // 替换为你的CA证书路径 .crt
		// 			ServerCertFile: "./server_ca.crt", // 替换为你的服务器证书路径 .crt
		// 			ServerKeyFile:  "./server_ca.key", // 替换为你的服务器私钥路径	.key
		// 		},
		// 	}
		// 	StartServer(httpsCACertConfig, dataStore, mu)
		// }()

		go func() {
			// HTTPS监听（自签名证书）
			httpServerAddr := ":8082"
			httpsSelfSignedConfig := &ServerConfig{
				Address: httpServerAddr,
				TLSConfig: &TLSConfig{
					ServerCertFile: "./server.crt", // 替换为你的自签名服务器证书路径 .crt
					ServerKeyFile:  "./server.key", // 替换为你的自签名服务器私钥路径 .key
				},
			}
			StartServer(httpsSelfSignedConfig, dataStore, mu)
		}()
	}

	fmt.Println("server is running...")

	select {
	case sig := <-sigChan:
		fmt.Println("sig:", sig)
		fmt.Println("Received exit signal. Performing cleanup...")
		// 在这里执行清理操作，例如关闭数据库连接等
	}
}

// 生成唯一的 orderid
func generateOrderID() string {
	// 在实际应用中，可以根据需要生成唯一的 orderid
	// 这里使用示例方法
	return uuid.New().String()
}

// StartServer 启动服务器根据提供的配置
func StartServer(config *ServerConfig, dataStore map[string]Data, mu *MutexWrapper) {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		handlePostRequest(w, r, dataStore, mu)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		handleGetRequest(w, r, dataStore, mu)
	})

	http.HandleFunc("/health", handleHealthCheck)

	if config.TLSConfig != nil {
		// HTTPS监听
		log.Printf("Server is running on %s (HTTPS)", config.Address)
		tlsConfig := &tls.Config{
			RootCAs:    x509.NewCertPool(),
			ServerName: "your-server-name", // 替换为您的服务器名称
		}

		if config.TLSConfig.CACertFile != "" {
			caCert, err := os.ReadFile(config.TLSConfig.CACertFile)
			if err != nil {
				log.Fatalf("Failed to read CA certificate: %v", err)
				return
			}
			if !tlsConfig.RootCAs.AppendCertsFromPEM(caCert) {
				log.Fatal("Failed to add CA certificate to the root CAs")
				return
			}
		}

		server := &http.Server{
			Addr:      config.Address,
			Handler:   nil, // 使用默认处理程序
			TLSConfig: tlsConfig,
		}
		err := server.ListenAndServeTLS(config.TLSConfig.ServerCertFile, config.TLSConfig.ServerKeyFile)
		if err != nil {
			log.Fatalf("Server error: %v", err)
		}
	} else {
		// HTTP监听
		log.Printf("Server is running on %s (HTTP)", config.Address)
		err := http.ListenAndServe(config.Address, nil)
		if err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}
}

// 处理 POST 请求
func handlePostRequest(w http.ResponseWriter, r *http.Request, dataStore map[string]Data, mu *MutexWrapper) {
	// POST 请求处理逻辑
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析 JSON 数据
	var jsonData Data
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// 检查参数是否为空
	if jsonData.ServerID == "" || jsonData.ValueID == "" || jsonData.AccountID == "" {
		http.Error(w, "serverid, valueid, and accountid must be provided", http.StatusBadRequest)
		return
	}

	// 生成一个唯一的 orderid
	orderID := generateOrderID()

	// 将 orderid 添加到 jsonData
	jsonData.OrderID = orderID

	// 存储 accountID 和 valueid 的映射关系
	mu.mu.Lock()
	dataStore[jsonData.OrderID] = jsonData
	mu.mu.Unlock()

	// 返回成功响应，包含 orderid
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonData)
}

// 处理 GET 请求
func handleGetRequest(w http.ResponseWriter, r *http.Request, dataStore map[string]Data, mu *MutexWrapper) {
	// GET 请求处理逻辑
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求参数
	orderID := r.FormValue("orderid")

	// 检查参数是否为空
	if orderID == "" {
		http.Error(w, "accountid must be provided", http.StatusBadRequest)
		return
	}

	// 查询 accountID 对应的值是否存在
	var data Data
	var exists bool
	mu.mu.Lock()
	data, exists = dataStore[orderID]
	mu.mu.Unlock()

	// 返回布尔值响应
	if exists {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "true")
		fmt.Println(data)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "false")
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	// 健康检查处理逻辑
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
