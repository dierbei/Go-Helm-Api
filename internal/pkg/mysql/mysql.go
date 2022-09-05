package mysql

import (
	"fmt"
	"log"
	"os"
	"sync"

	"githup.com/dierbei/go-helm-api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDb() *gorm.DB {
	once.Do(func() {
		c := config.GetConfig()
		_db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
				c.Mysql.Username,
				c.Mysql.Password,
				c.Mysql.Address,
				c.Mysql.Port,
				c.Mysql.Database,
			),
			//DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
			//DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			//DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			//DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			//SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		// 开发环境使用 gorm 默认日志
		// 生成环境禁用日志打印
		if os.Getenv(c.Application.Mode) != "debug" {
			_db.Logger = logger.Default.LogMode(logger.Silent)
		}

		db = _db
	})

	return db
}
