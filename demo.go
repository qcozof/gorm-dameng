/**
 * @Description $
 * @Author $
 * @Date $ $
 **/
package main

import (
	"fmt"
	"github.com/qcozof/gorm-dameng/dameng"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func main()  {

	fmt.Println("连接dameng...")
	gormConfig := config(true)
	GORM_DB, err := gorm.Open(dameng.Open( "dm://sysdba:dameng123!@193.100.100.221:5236?autoCommit=true"), gormConfig)

	if err != nil {
		fmt.Println("连接失败：", err)
		os.Exit(0)
	}
	fmt.Println("连接成功：",GORM_DB)
}



func config(logMode bool) (c *gorm.Config) {
	if logMode {
		c = &gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: 1 * time.Millisecond,
				LogLevel:      logger.Warn,
				Colorful:      true,
			}),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:"DW.",
				SingularTable: true, //表名后面不加s
			},

		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:"DW.",
				SingularTable: true, //表名后面不加s

			},
			//Namer.ColumnName : func() {
			//
			//},
		}
	}
	return
}