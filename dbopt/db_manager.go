package dbopt

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shanyongsy/go-web-test/msg"
)

type DBManager struct {
	db *sql.DB
	mu sync.Mutex
}

// NewDBManager 创建一个新的 DBManager 实例
func NewDBManager(host string, port string, user string, pwd string, dbname string) (*DBManager, error) {
	dsn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 设置连接池大小和最大闲置连接数
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// 设置连接超时时间
	db.SetConnMaxLifetime(time.Minute * 5)

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")

	return &DBManager{db: db}, nil
}

// Close 关闭数据库连接
func (m *DBManager) Close() {
	if m.db != nil {
		m.db.Close()
	}
}

func (m *DBManager) InsertRechargeRecord(record *msg.RechargeRequest, tradeNo string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 判断是否具有相同 ShopOrderNumber 的记录
	can, err := m.canInsert(record.OrderNo)
	if err != nil {
		return err
	}

	if !can {
		log.Printf("same shop order number: %v", record.OrderNo)
		return nil
	}

	// 向 DB 中插入 RechargeRecord 信息
	if err := m.insertRechargeRecord(record, tradeNo); err != nil {
		return err
	}

	return nil
}

// 是否已经发货
func (m *DBManager) HaveRealRecharges(tradeNo string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 根据 tradeNo 查询 RechargeRecord
	status, err := m.queryStatusByTradeNumber(tradeNo)

	if err != nil {
		return false, err
	}

	return status == 3, nil
}
