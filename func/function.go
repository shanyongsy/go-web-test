package function

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func PreStart() {
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// 在当前工作目录下创建 ./log 目录
	dir := filepath.Join(currentDir, "log")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 如果目录不存在，则创建
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}
}

// 生成唯一的 orderid
func GenerateTradeNo() string {
	// 在实际应用中，可以根据需要生成唯一的 orderid
	// 这里使用示例方法
	return uuid.New().String()
}

func StructValuesToString(s interface{}) string {
	// 获取结构体的反射值
	val := reflect.ValueOf(s)

	// 获取结构体的反射类型
	typ := val.Type()

	// 定义一个字符串变量用于保存结果
	result := ""

	// 遍历结构体的字段
	for i := 0; i < val.NumField(); i++ {
		// 获取字段的反射值
		fieldVal := val.Field(i)

		// 获取字段的名称
		fieldName := typ.Field(i).Name

		// 格式化为字符串并追加到结果中
		result += fmt.Sprintf("%s: %v\n", fieldName, fieldVal.Interface())
	}

	return result
}

func ParseAndLocalizeTime(timeStampStr string) (time.Time, error) {
	// 获取本地时区
	localLocation, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, fmt.Errorf("error loading local location: %v", err)
	}

	// 解析时间戳字符串，指定本地时区
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStampStr, localLocation)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %v", err)
	}

	return parsedTime, nil
}
