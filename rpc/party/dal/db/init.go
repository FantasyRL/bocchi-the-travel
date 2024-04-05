package db

import (
	"bocchi/pkg/constants"
	"bocchi/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DBMember *gorm.DB
	DBParty  *gorm.DB
)

func Init() {
	var err error
	DBParty, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("mysql connect error")
	} else {
		klog.Info("mysql connect access")
	}

	sqlDB, err := DBParty.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBParty = DBParty.Table(constants.PartyTableName)

	DBMember, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("mysql:2 connect error")
	} else {
		klog.Info("mysql:2 connect access")
	}
	sqlDB, err = DBMember.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBMember = DBMember.Table(constants.MemberTableName)
}
