package gorm

import (
	"fmt"
	"log"
	"sync"

	"github.com/shanyongsy/go-web-test/model"
	"github.com/shanyongsy/go-web-test/msg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBManager struct {
	db *gorm.DB
	mu sync.Mutex
}

// NewDBManager 创建一个新的 DBManager 实例
func NewDBManager(host string, port string, user string, pwd string, dbname string) (*DBManager, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("[DB] error connecting to database: %v", err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&model.RechargeInfo{})
	if err != nil {
		log.Fatalf("[DB] Error migrating database: %v", err)
	}

	err = db.AutoMigrate(&model.SimpleRechargeInfo{})
	if err != nil {
		log.Fatalf("[DB] Error migrating database: %v", err)
	}

	return &DBManager{db: db}, nil
}

// Close 关闭数据库连接
func (m *DBManager) Close() {
}

func (m *DBManager) InsertRechargeRecord(record *msg.RechargeRequest, tradeNo string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 判断是否具有相同 ShopOrderNumber 的记录
	can, err := m.canInsertTheOrder(record.OrderNo)
	if err != nil {
		return false, err
	}

	if can == false {
		log.Printf("[DB] Have same shop order, number  is %v, can't insert.", record.OrderNo)
		return false, nil
	}

	// 向 DB 中插入 RechargeRecord 信息
	if err := m.insertRechargeRecord(record, tradeNo); err != nil {
		log.Printf("[DB] Failed to insert recharge, record is %v , err is %v", record, err)
		return false, err
	}
	log.Printf("[DB] Insert recharge success, order is %v, tradeNo is %v", record.OrderNo, tradeNo)
	return true, nil
}

// 是否已经发货完毕
func (m *DBManager) HaveFinishRecharges(tradeNo string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 根据 tradeNo 查询 RechargeRecord
	status, err := m.queryStatusByTradeNumber(tradeNo)

	if err != nil {
		return false, err
	}

	return status == 3, nil
}

// 改变发货状态
func (m *DBManager) ChangeRechargeStatus(tradeNo string, status int32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 根据 tradeNo 查询 RechargeRecord
	err := m.updateStatusByTradeNumber(tradeNo, status)

	if err != nil {
		return err
	}

	return nil
}

// 获取一条尚未发货的记录
func (m *DBManager) GetOneUnfinishedRecharge() (*model.RechargeInfo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	have, err := m.haveUnfinishedRecharge()
	if err != nil {
		return nil, err
	}
	if have == false {
		return nil, nil
	}

	// 根据 tradeNo 查询 RechargeRecord
	rechargeInfo, err := m.getOneUnfinishedRecharge()

	if err != nil {
		return nil, err
	}

	return rechargeInfo, nil
}
