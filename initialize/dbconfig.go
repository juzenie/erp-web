package initialize

import (
	"encoding/json"
	"erp-web/config"
	"erp-web/global"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func Initdb() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic("读取数据库配置异常")
	}
	mysqlInfo := &config.Mysql{}
	if err := v.Unmarshal(mysqlInfo); err != nil {
		panic("<empty>")
	}
	fmt.Println(mysqlInfo)
	data, _ := json.Marshal(mysqlInfo)
	zap.S().Info(string(data))

	/*	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			dbInfo.Host, dbInfo.User, dbInfo.Password, dbInfo.Dbname, dbInfo.Port,
		)
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 禁用彩色打印
			},
		)
		db, err := gorm.Open(pg.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
	*/
	dbInfo := mysqlInfo.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbInfo.User, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.Dbname)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
