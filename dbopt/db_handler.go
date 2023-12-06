package dbopt

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/shanyongsy/go-web-test/model"
	"github.com/shanyongsy/go-web-test/msg"
)

// 向 DB 中插入 RechargeRecord 信息
func (m *DBManager) insertRechargeRecord(record *msg.RechargeRequest, tradeNo string) error {

	// TradeNumber := tradeNo
	// ShopOrderNumber := record.OrderNo
	// ShopType := int32(0)
	// {
	// 	i, err := strconv.Atoi(record.ShopType)
	// 	if err != nil {
	// 		log.Printf("error converting shop type to int: %v, str is %v", err, record.ShopType)
	// 	}
	// 	ShopType = int32(i)
	// }
	// ShopGoodsID := record.GoodsID
	// ShopGoodsName := record.GoodsName
	// ShopAccountID := record.AccountID
	// ShopBuyerPhoneNumber := record.BuyerPhoneNumber
	// ShopBuyerID := record.BuyerID
	// ShopOrderCreateAt := time.Now()
	// {
	// 	t, err := time.Parse("2006-01-02 15:04:05", record.TimeStamp)
	// 	if err != nil {
	// 		log.Printf("error converting timestamp to time.Time: %v, str is %v", err, record.TimeStamp)
	// 	}
	// 	ShopOrderCreateAt = t
	// }
	// Amount := 0.0
	// {
	// 	f, err := strconv.ParseFloat(record.Amount, 64)
	// 	if err != nil {
	// 		log.Printf("error converting amount to float64: %v, str is %v", err, record.Amount)
	// 	}
	// 	Amount = f
	// }
	// SingleAmount := 0.0
	// {
	// 	f, err := strconv.ParseFloat(record.SingleAmount, 64)
	// 	if err != nil {
	// 		log.Printf("error converting single amount to float64: %v, str is %v", err, record.SingleAmount)
	// 	}
	// 	SingleAmount = f
	// }
	// TotalAmount := 0.0
	// {
	// 	f, err := strconv.ParseFloat(record.TotalAmount, 64)
	// 	if err != nil {
	// 		log.Printf("error converting total amount to float64: %v, str is %v", err, record.TotalAmount)
	// 	}
	// 	TotalAmount = f
	// }
	// Count := 0
	// {
	// 	i, err := strconv.Atoi(record.Count)
	// 	if err != nil {
	// 		log.Printf("error converting count to int: %v, str is %v", err, record.Count)
	// 	}
	// 	Count = int32(i)
	// }
	// RealRechargeCount := 0
	// TradeCreateAt := time.Now()
	// TradeUpdateAt := time.Now()
	// TryRechargeCount := 0
	// Status := 0
	// GameMoney := 0
	// GameType := 0
	// {
	// 	i, err := strconv.Atoi(record.GameType)
	// 	if err != nil {
	// 		log.Printf("error converting game type to int: %v, str is %v", err, record.GameType)
	// 	}
	// 	GameType = int32(i)
	// }

	// sql := "insert into recharge_info(
	// 	trade_number,
	// 	shop_order_number,
	// 	shop_type,
	// 	shop_goods_id,
	// 	shop_goods_name,
	// 	shop_account_id,
	// 	shop_buyer_phone_number,
	// 	shop_buyer_id,
	// 	shop_order_create_at,
	// 	amount, single_amount,
	// 	total_amount,
	// 	count,
	// 	real_recharge_count,
	// 	trade_create_at,
	// 	trade_update_at,
	// 	try_recharge_c

	return nil
}

// 判断是否具有相同 ShopOrderNumber 的记录
func (m *DBManager) canInsert(shopOrderNumber string) (bool, error) {

	row := m.db.QueryRow("select id from recharge_info where shop_order_number = ?", shopOrderNumber)
	if row.Err() == sql.ErrNoRows {
		return true, nil
	}

	if row.Err() != nil {
		return false, row.Err()
	}

	return false, nil
}

// 查询数据表 recharge_info 中指定 TradeNumber 的 status 字段
func (m *DBManager) queryStatusByTradeNumber(tradeNumber string) (int32, error) {

	row := m.db.QueryRow("select status from recharge_info where trade_number = ?", tradeNumber)
	if row.Err() == sql.ErrNoRows {
		return 0, nil
	}

	if row.Err() != nil {
		return 0, row.Err()
	}

	var status int32

	err := row.Scan(&status)
	if err != nil {
		return 0, err
	}

	return status, nil
}

// 更新数据表 recharge_info 中指定 ID 的 status 字段
func (m *DBManager) updateStatusByTradeNumber(tradeNumber string, status int32) error {

	return nil
}

// 获取指定 status 的记录
func (m *DBManager) getRechargeInfoByStatus(status int32) ([]model.RechargeInfo, error) {

	return nil, nil
}
