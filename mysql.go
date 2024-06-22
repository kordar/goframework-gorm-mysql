package goframework_gorm_mysql

import (
	"github.com/kordar/godb"
	log "github.com/kordar/gologger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mysqlpool  = godb.NewDbPool()
	dbLogLevel = "info"
)

func SetDbLogLevel(level string) {
	dbLogLevel = level
}

func GetMysqlDB(db string) *gorm.DB {
	return mysqlpool.Handle(db).(*gorm.DB)
}

func gormConfig() *gorm.Config {
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
func InitMysqlHandle(dbs map[string]map[string]string) {
	for db, cfg := range dbs {
		ins := NewGormConnIns(db, cfg, gormConfig())
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
func AddMysqlInstance(db string, cfg map[string]string) error {
	ins := NewGormConnIns(db, cfg, gormConfig())
	return mysqlpool.Add(ins)
}

func AddMysqlInstanceWithDsn(db string, dsn string) error {
	ins := NewGormConnInsWithDsn(db, dsn, gormConfig())
	return mysqlpool.Add(ins)
}

// RemoveMysqlInstance 移除mysql句柄
func RemoveMysqlInstance(db string) {
	mysqlpool.Remove(db)
}

// HasMysqlInstance mysql句柄是否存在
func HasMysqlInstance(db string) bool {
	return mysqlpool != nil && mysqlpool.Has(db)
}
