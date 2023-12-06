package msg

// Recharge request
type RechargeRequestWithSpider struct {
	TradeNo   string `json:"tradeNo" binding:"required"`   // 订单编号
	GoodsID   string `json:"goodsID" binding:"required"`   // 商品编号
	GoodsName string `json:"goodsName" binding:"required"` // 商品名称
	AccountID string `json:"accountID" binding:"required"` // 充值账号
	Count     int32  `json:"count" binding:"required"`     // 数量
	GameType  int32  `json:"gameType" binding:"required"`  // 区服 1:通宝区 2:元宝区
	ShopType  int32  `json:"shopType" binding:"required"`  // 商城类型 1:淘宝 2:京东 3:拼多多
	GameMoney int32  `json:"gameMoney" binding:"required"` // 游戏币
}

// Recharge response
type RechargeResponseWithSpider struct {
	Status bool `json:"status" binding:"required"` // 状态 true:成功 false:失败
}
