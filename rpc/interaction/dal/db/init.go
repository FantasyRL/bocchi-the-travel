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
	DBComment *gorm.DB
	//DBCommentLike *gorm.DB
)

func Init() {
	var err error
	DBComment, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("db connect error")
	} else {
		klog.Info("db connect access")
	}
	sqlDB, _ := DBComment.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBComment = DBComment.Table(constants.CommentTableName).WithContext(context.Background())
}
