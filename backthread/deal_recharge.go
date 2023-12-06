package backthread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	gorm_mgr "github.com/shanyongsy/go-web-test/gorm"
	"github.com/shanyongsy/go-web-test/model"
	"github.com/shanyongsy/go-web-test/msg"
)

func ProcessRecharge(dbManager *gorm_mgr.DBManager, exitChan <-chan struct{}) {
	for {
		select {
		case <-exitChan:
			log.Println("[back] Received exit signal. Cleaning up and exiting...")
			// 执行清理操作，例如关闭资源等
			return
		default:
			// 模拟后台处理逻辑，每隔一段时间处理一次
			time.Sleep(1 * time.Millisecond)
			// log.Println("[back] Processing ProcessRecharge data from DB...")
			// 添加你的实际处理逻辑

			// 从数据库中读取尚未进行充值的记录
			info, err := dbManager.GetOneUnfinishedRecharge()
			if err != nil {
				log.Printf("[back] Getting unfinished recharge, error: %v", err)
				continue
			}

			if info == nil { // 没有未完成的充值记录
				time.Sleep(1 * time.Second)
				continue
			}

			// 向爬虫系统发送充值请求
			// sus, err := RechargeWithSpider(info)
			// time.Sleep(1 * time.Millisecond)
			sus := true

			// 更新数据库中的记录
			if sus {
				err = dbManager.ChangeRechargeStatus(info.TradeNumber, 3)
				if err != nil {
					log.Printf("[back] error changing recharge status: %v, id is %v, trade is %v", err, info.ID, info.TradeNumber)
					continue
				} else {
					log.Printf("[back] change recharge status to 3, id is %v, trade is %v", info.ID, info.TradeNumber)
				}
			}
		}
	}
}

func rechargeWithSpider(info *model.RechargeInfo) (bool, error) {

	client := &http.Client{}

	request := msg.RechargeRequestWithSpider{}
	request.TradeNo = info.TradeNumber
	request.GoodsID = info.ShopGoodsID
	request.GoodsName = info.ShopGoodsName
	request.AccountID = info.ShopAccountID
	request.Count = info.Count
	request.GameType = info.GameType
	request.ShopType = info.ShopType
	request.PriceType = info.PriceType

	requestJSON, _ := json.Marshal(request)
	postURL := "http://localhost:8080/spider"

	resp, err := client.Post(postURL, "application/json; charset=utf-8", bytes.NewBuffer(requestJSON))
	if err != nil {
		fmt.Println("POST request error:", err)
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("POST request failed with status code:", resp.Status)
		return false, nil
	}

	responseJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read POST response:", err)
		return false, err
	}

	responseInfo := msg.RechargeResponseWithSpider{}
	err = json.Unmarshal(responseJSON, &responseInfo)
	if err != nil {
		fmt.Println("Failed to unmarshal POST response data:", err)
		return false, err
	}

	return responseInfo.Status, nil
}
