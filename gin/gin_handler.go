package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanyongsy/go-web-test/msg"
)

func handleInterRecharge(c *gin.Context) {
	var request msg.RechargeRequest

	log.Println("[HTTP] inter recharge request 1:", c.Request)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("[HTTP] inter recharge request 2:", request)

	// tradeNo := generateTradeNo()
	var tradeNo string
	if request.ShopType == "1" {
		tradeNo = "taobao_" + request.OrderNo
	} else if request.ShopType == "2" {
		tradeNo = "tianmao_" + request.OrderNo
	} else if request.ShopType == "3" {
		tradeNo = "jingdong_" + request.OrderNo
	} else {
		c.JSON(http.StatusOK, &msg.RechargeResponse{Status: false, TradeNo: request.OrderNo})
		return
	}

	var sus bool = true
	// var err error

	// // 存入数据库
	// db := c.MustGet("db").(*gorm_mgr.DBManager)
	// if sus, err = db.InsertRechargeRecord(&request, tradeNo); err != nil {
	// 	log.Printf("[HTTP] inter recharge error: %v, request is %v", err.Error(), request)
	// }

	var response msg.RechargeResponse
	response.Status = sus
	response.TradeNo = tradeNo

	log.Println("[HTTP] inter recharge response:", response)

	c.JSON(http.StatusOK, response)
}

func handleInterCheck(c *gin.Context) {
	var request msg.RechargeResultRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// log.Println("check request:", request)

	var response msg.RechargeResultResponse
	response.Status = false

	// log.Println("check response:", response)

	// 根据 tradeNo 查询 status
	// db := c.MustGet("db").(*gorm_mgr.DBManager)
	// haveRecharge, err := db.HaveFinishRecharges(request.TradeNo)
	// if err != nil {
	// 	log.Printf("[HTTP] check error: %v, request is %v", err.Error(), request)
	// } else {
	// 	response.Status = haveRecharge
	// }
	// log.Printf("[HTTP] check request is %v, response: %v", request, response)

	c.JSON(http.StatusOK, response)
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")

	log.Println("pong")
}

func handleInterSimpleRecharge(c *gin.Context) {
	var request msg.RechargeRequest

	info := "[HTTP] inter simple recharge "

	log.Println(info+"request 1:", c.Request)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(info+"request 2:", request)

	c.JSON(http.StatusOK, gin.H{"status": "true"})
}

func handelChangeStatus(c *gin.Context) {
	var request msg.ChangeRechargeStatusRequest
	var response msg.ChangeRechargeStatusResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Status = false
		c.JSON(http.StatusOK, response)
		return
	}

	// db := c.MustGet("db").(*gorm_mgr.DBManager)
	// if err := db.ChangeRechargeStatus(request.TradeNo, request.Status); err != nil {
	// 	response.Status = false
	// 	c.JSON(http.StatusOK, response)
	// 	return
	// }

	response.Status = true
	c.JSON(http.StatusOK, response)
}

// 模拟爬虫系统的接口
func handleInterRechargeFromSpider(c *gin.Context) {
	var request msg.RechargeRequestWithSpider
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("[HTTP][Spider], recharge request, value is ", request)

	var response msg.RechargeResponseWithSpider
	response.Status = true

	// 人为设置延迟
	// time.Sleep(16 * time.Second)

	log.Println("[HTTP][Spider], recharge response, value is ", response)
	c.JSON(http.StatusOK, response)
}
