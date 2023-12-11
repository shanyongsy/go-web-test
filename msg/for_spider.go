package msg

// Recharge request
type RechargeRequestWithSpider struct {
	TradeNo     string `json:"tradeNo" validate:"required" msg:"缺少订单编号"`       // 订单编号(发货服务数据，根据商城订单编号生成的订单编号)
	ShopOrderNo string `json:"shopOrderNo" validate:"required" msg:"缺少商城订单编号"` // 商城订单编号（商城数据，商城：淘宝、等）
	GoodsID     string `json:"goodsID" validate:"required" msg:"缺少商品编号"`       // 商品编号（商城数据）
	GoodsName   string `json:"goodsName" validate:"required" msg:"缺少商品名称"`     // 商品名称（商城数据）
	ShopType    int32  `json:"shopType" validate:"required" msg:"缺少商城类型"`      // 商城类型 1:淘宝 2:京东 3:拼多多
	ChargeGame  int32  `json:"chargeGame" validate:"required" msg:"缺少充值游戏产品"`  // 产品 充值游戏产品 (1:元宝区（封神榜） 2:通宝区（封神榜国际版）)
	GameAccount string `json:"gameAccount" validate:"required" msg:"缺少玩家游戏账号"` // 充值账号（商城数据，即 玩家游戏账号、确认游戏账号）
	CardType    int32  `json:"cardType" validate:"required" msg:"缺少充入卡类型"`     // 付费卡类型（即 充入卡类型，(61：5元 62：15元 63：50元 65：100元））
	CardAmount  int32  `json:"cardAmount" validate:"required" msg:"缺少充入卡数量"`   // 数量（商城数据，即 充入卡数量）
	ChargeMode  int32  `json:"chargeMode" validate:"required" msg:"缺少充入卡方式"`   // 充入卡方式 1：包时 2：计点 3：元宝 5：封神通宝
}

// Recharge response
type RechargeResponseWithSpider struct {
	Message string `json:"message"`                   // 消息
	Status  bool   `json:"status" binding:"required"` // 状态 true:成功 false:失败
}
