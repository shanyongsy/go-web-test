package msg

// Recharge request
type RechargeRequest struct {
	OrderNo          string `json:"orderNo" binding:"required"`      // 订单编号
	GoodsID          string `json:"goodsID" binding:"required"`      // 商品编号
	GoodsName        string `json:"goodsName" binding:"required"`    // 商品名称
	AccountID        string `json:"accountID" binding:"required"`    // 充值账号
	BuyerPhoneNumber string `json:"buyerPhoneNumber"`                // 买家手机号
	BuyerID          string `json:"buyerID" binding:"required"`      // 买家ID
	Amount           string `json:"amount" binding:"required"`       // 金额
	SingleAmount     string `json:"singleAmount" binding:"required"` // 单价
	TotalAmount      string `json:"totalAmount" binding:"required"`  // 总价
	Count            string `json:"count" binding:"required"`        // 数量
	TimeStamp        string `json:"timeStamp" binding:"required"`    // 时间戳
	GameType         string `json:"gameType" binding:"required"`     // 区服 1:通宝区 2:元宝区
	ShopType         string `json:"shopType" binding:"required"`     // 商城类型 1:淘宝 2:京东 3:拼多多
	PriceType        string `json:"priceType" binding:"required"`    // 价格类型 15 30 50 100
}

// Recharge response
type RechargeResponse struct {
	Status  bool   `json:"status" binding:"required"`  // 状态 true:成功 false:失败
	TradeNo string `json:"tradeNo" binding:"required"` // 交易号，本服务生成
}

// Simple recharge request
type SimpleRechargeRequest struct {
	OrderNo          string `json:"orderNo" binding:"required"`   // 订单编号
	GoodsID          string `json:"goodsID" binding:"required"`   // 商品编号
	GoodsName        string `json:"goodsName" binding:"required"` // 商品名称
	AccountID        string `json:"accountID" binding:"required"` // 充值账号
	BuyerPhoneNumber string `json:"buyerPhoneNumber"`             // 买家手机号
	BuyerID          string `json:"buyerID" binding:"required"`   // 买家ID
}

type SimpleRechargeResponse struct {
	Status string `json:"status" binding:"required"` // 状态 true:成功 false:失败
}

// Recharge result request
type RechargeResultRequest struct {
	TradeNo string `json:"tradeNo" binding:"required"` // 交易号，本服务生成
}

// Recharge result response
type RechargeResultResponse struct {
	Status bool `json:"status" binding:"required"` // 状态 true:成功 false:失败
}

type ChangeRechargeStatusRequest struct {
	TradeNo string `json:"tradeNo" binding:"required"` // 交易号，本服务生成
	Status  int32  `json:"status" binding:"required"`  // 状态 0-未处理 3-已完成
}

type ChangeRechargeStatusResponse struct {
	Status bool `json:"status" binding:"required"` // 状态 true:成功 false:失败
}
