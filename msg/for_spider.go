package msg

// Recharge request
type RechargeRequestWithSpider struct {
	TradeNo     string `json:"tradeNo" binding:"required"`         // 订单编号(发货服务数据，根据商城订单编号生成的订单编号)
	ShopOrderNo string `json:"shopOrderNo" binding:"required"`     // 商城订单编号（商城数据，商城：淘宝、等）
	GoodsID     string `json:"goodsID" binding:"required"`         // 商品编号（商城数据）
	GoodsName   string `json:"goodsName" binding:"required"`       // 商品名称（商城数据）
	AccountID   string `json:"accountID" binding:"required"`       // 充值账号（商城数据，即 玩家游戏账号、确认游戏账号）
	Count       int32  `json:"count" binding:"required"`           // 数量（商城数据，即 充入卡数量）
	GameType    int32  `json:"gameType" binding:"required"`        // 产品（即 充值游戏产品 1:通宝区（封神榜国际版） 2:元宝区（封神榜））
	ShopType    int32  `json:"shopType" binding:"required"`        // 商城类型 1:淘宝 2:京东 3:拼多多
	PriceType   int32  `json:"PriceType" binding:"required"`       // 付费卡类型（即 充入卡类型，取值为：5、15、50、100）
	ValueType   int32  `json:"entryGameMethod" binding:"required"` // 充入卡方式 101:封神通宝 201：包时 202：计点 203：元宝
}

// Recharge response
type RechargeResponseWithSpider struct {
	Status bool `json:"status" binding:"required"` // 状态 true:成功 false:失败
}
