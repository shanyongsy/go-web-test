# 生成 models
### 下载工具
```bash
go install gorm.io/gen/tools/gentool@latest
```
### 添加执行目录
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
### 修改配置
修改 gen.sh 以便决定生成哪些 models
```bash
vim gen.sh
```
### 生成 models
```bash
./gen.sh
```
### 修改 .gen.go
对于自动迁移的表，如果字符串型字段是索引，那么需要修改相应的 .gen.go 文件
```bash
type:varchar(255);
```
