package gorm

import (
	"errors"
	"log"
	"strconv"
	"time"

	function "github.com/shanyongsy/go-web-test/func"
	"github.com/shanyongsy/go-web-test/model"
	"github.com/shanyongsy/go-web-test/msg"
)

// 向 DB 中插入 RechargeRecord 信息
func (m *DBManager) insertRechargeRecord(record *msg.RechargeRequest, tradeNo string) error {

	// 将 RechargeRequest 转换为 RechargeInfo
	rechargeInfo := model.RechargeInfo{}

	rechargeInfo.TradeNumber = tradeNo
	rechargeInfo.ShopOrderNumber = record.OrderNo
	{
		i, err := strconv.Atoi(record.ShopType)
		if err != nil {
			log.Printf("[DB] error converting shop type to int: %v, str is %v", err, record.ShopType)
		}
		rechargeInfo.ShopType = int32(i)
	}
	rechargeInfo.ShopGoodsID = record.GoodsID
	rechargeInfo.ShopGoodsName = record.GoodsName
	rechargeInfo.ShopAccountID = record.AccountID
	rechargeInfo.ShopBuyerPhoneNumber = record.BuyerPhoneNumber
	rechargeInfo.ShopBuyerID = record.BuyerID
	{

		t, err := function.ParseAndLocalizeTime(record.TimeStamp)
		if err != nil {
			log.Printf("[DB] error converting timestamp to time.Time: %v, str is %v", err, record.TimeStamp)
		}
		rechargeInfo.ShopOrderCreateAt = t.Local()
	}
	{
		f, err := strconv.ParseFloat(record.Amount, 64)
		if err != nil {
			log.Printf("[DB] error converting amount to float64: %v, str is %v", err, record.Amount)
		}
		rechargeInfo.Amount = f
	}
	{
		f, err := strconv.ParseFloat(record.SingleAmount, 64)
		if err != nil {
			log.Printf("[DB] error converting single amount to float64: %v, str is %v", err, record.SingleAmount)
		}
		rechargeInfo.SingleAmount = f
	}
	{
		f, err := strconv.ParseFloat(record.TotalAmount, 64)
		if err != nil {
			log.Printf("[DB] error converting total amount to float64: %v, str is %v", err, record.TotalAmount)
		}
		rechargeInfo.TotalAmount = f
	}
	{
		i, err := strconv.Atoi(record.Count)
		if err != nil {
			log.Printf("[DB] error converting count to int: %v, str is %v", err, record.Count)
		}
		rechargeInfo.Count = int32(i)
	}
	rechargeInfo.RealRechargeCount = 0
	rechargeInfo.TradeCreateAt = time.Now()
	rechargeInfo.TradeUpdateAt = time.Now()
	rechargeInfo.TryRechargeCount = 0
	rechargeInfo.Status = 0
	{
		i, err := strconv.Atoi(record.PriceType)
		if err != nil {
			log.Printf("[DB] error converting game money to int: %v, str is %v", err, record.PriceType)
		}
		rechargeInfo.GameMoney = int32(i)
	}
	{
		i, err := strconv.Atoi(record.GameType)
		if err != nil {
			log.Printf("[DB] error converting game type to int: %v, str is %v", err, record.GameType)
		}
		rechargeInfo.GameType = int32(i)
	}

	// log.Printf("[DB] Insert value , TradeCreateAt=%v, TradeUpdateAt=%v, ShopOrderCreateAt=%v, TimeStamp=%v",
	// 	rechargeInfo.TradeCreateAt,
	// 	rechargeInfo.TradeUpdateAt,
	// 	rechargeInfo.ShopOrderCreateAt,
	// 	record.TimeStamp)

	// 合理性检查
	{
		// ShopType
		switch rechargeInfo.ShopType {
		case 1: // 淘宝
		case 2: // 天猫
		case 3: // 京东
		case 4: // 拼多多
			{
				break
			}
		default:
			{
				log.Printf("[DB] ShopType is invalid, ShopType is %v", rechargeInfo.ShopType)
				return errors.New("ShopType is invalid")
			}
		}

		// GameType
		switch rechargeInfo.GameType {
		case 1: // 通宝区
		case 2: // 元宝区
			{
				break
			}
		default:
			{
				log.Printf("[DB] GameType is invalid, GameType is %v", rechargeInfo.GameType)
				return errors.New("GameType is invalid")
			}
		}

		// GameMoney
		switch rechargeInfo.GameMoney {
		case 1: // 1
		case 15: // 15
		case 30: // 30
		case 50: // 50
		case 100: // 100
			{
				break
			}
		default:
			{
				log.Printf("[DB] GameMoney is invalid, GameMoney is %v", rechargeInfo.GameMoney)
				return errors.New("GameMoney is invalid")
			}
		}

		// Count
		if rechargeInfo.Count <= 0 || rechargeInfo.Count > 100 {
			log.Printf("[DB] Count is invalid, Count is %v", rechargeInfo.Count)
			return errors.New("Count is invalid")
		}
	}

	return m.db.Create(&rechargeInfo).Error
}

func (m *DBManager) insertSimpleRechargeRecord(record *msg.RechargeRequest, tradeNo string) error {
	// 将 RechargeRequest 转换为 SimpleRechargeInfo
	rechargeInfo := model.SimpleRechargeInfo{}

	rechargeInfo.TradeNumber = tradeNo
	rechargeInfo.Count = 1

	return m.db.Create(&rechargeInfo).Error
}

// 判断是否具有相同 ShopOrderNumber 的记录
func (m *DBManager) canInsertTheOrder(shopOrderNumber string) (bool, error) {
	var count int64
	resault := m.db.Model(&model.RechargeInfo{}).Where("shop_order_number = ?", shopOrderNumber).Count(&count)

	if resault.Error != nil {
		return false, resault.Error
	}

	log.Printf("[DB] OrderNo is %v, count is %v", shopOrderNumber, count)
	return count == 0, nil
}

// 查询数据表 recharge_info 中指定 TradeNumber 的 status 字段
func (m *DBManager) queryStatusByTradeNumber(tradeNumber string) (int32, error) {
	var rechargeInfo model.RechargeInfo
	resault := m.db.Where("trade_number = ?", tradeNumber).First(&rechargeInfo)
	if resault.Error != nil {
		log.Printf("[DB] queryStatusByTradeNumber error: %v, tradeNumber is %v", resault.Error, tradeNumber)
		return 0, resault.Error
	}

	log.Printf("[DB] TradeNo is %v, status is %v", tradeNumber, rechargeInfo.Status)
	return rechargeInfo.Status, nil
}

// 更新数据表 recharge_info 中指定 ID 的 status 字段
func (m *DBManager) updateStatusByTradeNumber(tradeNumber string, status int32) error {
	resault := m.db.Model(&model.RechargeInfo{}).
		Where("trade_number = ?", tradeNumber).
		Updates(map[string]interface{}{
			"status":          status,
			"trade_update_at": time.Now(),
		})
	if resault.Error != nil {
		log.Printf("[DB] updateStatusByTradeNumber error: %v, tradeNumber is %v, status is %v", resault.Error, tradeNumber, status)
		return resault.Error
	}

	log.Printf("[DB] updateStatusByTradeNumber rows affected is %v, tradeNumber is %v, status is %v", resault.RowsAffected, tradeNumber, status)
	return nil
}

// 获取指定 status 的记录
func (m *DBManager) getRechargeInfoByStatus(status int32) ([]model.RechargeInfo, error) {
	var rechargeInfos []model.RechargeInfo
	resault := m.db.Where("status = ?", status).Limit(100).Find(&rechargeInfos)
	if resault.Error != nil {
		log.Printf("[DB] getRechargeInfoByStatus error: %v, status is %v", resault.Error, status)
		return nil, resault.Error
	}

	log.Printf("[DB] getRechargeInfoByStatus rows affected is %v, status is %v", resault.RowsAffected, status)
	return rechargeInfos, nil
}

// 判断是否还有尚未充值完成的的记录
func (m *DBManager) haveUnfinishedRecharge() (bool, error) {
	var count int64
	resault := m.db.Model(&model.RechargeInfo{}).Where("status = ?", 0).Count(&count)

	if resault.Error != nil {
		return false, resault.Error
	}

	return count > 0, nil
}

func (m *DBManager) getOneUnfinishedRecharge() (*model.RechargeInfo, error) {
	var rechargeInfo model.RechargeInfo
	resault := m.db.Where("status = ?", 0).First(&rechargeInfo)
	if resault.Error != nil {
		log.Printf("[DB] Get one unfinished recharge, error: %v", resault.Error)
		return nil, resault.Error
	}

	log.Printf("[DB] Get one unfinished recharge, rows affected is %v, tradeNo is %v", resault.RowsAffected, rechargeInfo.TradeNumber)
	return &rechargeInfo, nil
}
