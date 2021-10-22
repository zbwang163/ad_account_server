package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Db          = make(map[string]*gorm.DB, 2)
	mysqlConfig = map[string]string{
		"ad.info.account_server": "root:123456@tcp(localhost:3306)/ad_account?charset=utf8mb4&parseTime=True&loc=Local",
		"ad.info.content_server": "root:123456@tcp(localhost:3306)/ad_content?charset=utf8mb4&parseTime=True&loc=Local",
	}
)

func InitMysql(psm string) {
	dsn, ok := mysqlConfig[psm]
	if !ok {
		return
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             0,           // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound biz_error for logger
			Colorful:                  false,       // Disable color
		},
	)
	// Globally mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("init mysql err:%v", err))
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	Db[psm] = db
}
