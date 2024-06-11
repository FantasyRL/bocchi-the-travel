package db

import (
	"bocchi/pkg/constants"
	"bocchi/pkg/utils"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DBFollow, DBMark, DBScore *gorm.DB
)

func Init() {
	var err error
	DBFollow, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
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
	sqlDB, _ := DBFollow.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBFollow = DBFollow.Table(constants.FollowTableName).WithContext(context.Background())

	DBMark, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("mysql mark connect error")
	} else {
		klog.Info("mysql mark connect access")
	}
	sqlDB, _ = DBMark.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBMark = DBMark.Table(constants.MarkTableName).WithContext(context.Background())

	DBScore, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("mysql score connect error")
	} else {
		klog.Info("mysql score connect access")
	}
	sqlDB, _ = DBScore.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBScore = DBScore.Table(constants.ScoreTableName).WithContext(context.Background())
}
