package goframework_gorm_mysql

import (
	"github.com/kordar/gocfg"
	"github.com/kordar/godb"
	log "github.com/kordar/gologger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlpool *godb.DbConnPool

func GetMysqlDB(db string) *gorm.DB {
	return mysqlpool.Handle(db).(*gorm.DB)
}

func gormConfig() *gorm.Config {
	dbLogLevel := gocfg.GetSystemValue("gorm_log_level")
	mysqlConfig := gorm.Config{}
	if dbLogLevel == "error" {
		mysqlConfig.Logger = logger.Default.LogMode(logger.Error)
	}
	if dbLogLevel == "warn" {
		mysqlConfig.Logger = logger.Default.LogMode(logger.Warn)
	}
	if dbLogLevel == "info" {
		mysqlConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	return &mysqlConfig
}

// InitMysqlHandle 初始化mysql句柄
func InitMysqlHandle(dbs ...string) {
	mysqlpool = godb.GetDbPool()
	for _, s := range dbs {
		ins := NewGormConnIns(s, gormConfig())
		if ins == nil {
			continue
		}
		err := mysqlpool.Add(ins)
		if err != nil {
			log.Warnf("初始化异常，err=%v", err)
		}
	}

}

// AddMysqlInstance 添加mysql句柄
func AddMysqlInstance(db string) error {
	mysqlpool = godb.GetDbPool()
	ins := NewGormConnIns(db, gormConfig())
	return mysqlpool.Add(ins)
}

// RemoveMysqlInstance 移除mysql句柄
func RemoveMysqlInstance(db string) {
	mysqlpool.Remove(db)
}