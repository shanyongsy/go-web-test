# go-web-test
### 环境
```bash
# 启动 MySQL
cd ./mysql
./start.sh
```
### 编译
```bash
./build.sh
```
### 启动 service
```bash
./run
```
#### run 命令行参数
```bash
$ ./run --help
Usage of ./run:
  -dbhost string
        database host (default "localhost")
  -dbname string
        database name (default "recharge_db")
  -dbport string
        database port (default "3306")
  -dbpwd string
        database password (default "123456")
  -dbuser string
        database user (default "root")
  -https
        use HTTPS
  -port string
        server port (default "8080")
```
### JSON 配置
```JSON
{"orderNo":"{$订单编号}","goodsID":"{$商品ID}","goodsName":"{$商品名称}","accountID":"{$订单扩展参数(游戏账号)}","buyerPhoneNumber":"{$联系人手机}","buyerID":"{$买家ID}","amount":"{$订单金额}","singleAmount":"{$商品实付单价}","totalAmount":"{$商品实付总价}","count":"{$购买数量}","timeStamp":"{$时间戳B}","gameType":"1","shopType":"1","priceType":"1"}
```
- gameType
    - 1：通宝区
    - 2：元宝区
- shopType
    - 1：淘宝
    - 2：天猫
    - 3：京东
    - 4：拼多多
- priceType
    - 1
    - 15
    - 30
    - 50
    - 100
### To do
- [x] DB 数据中加入尝试发货次数，超过次数不再发货
- [x] 后台线程处理时优先处理尝试次数少的数据，并且尝试次数加 1
- [x] 另起一个监听，模拟爬虫 service

# go-web-test
