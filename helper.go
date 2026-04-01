package goframework_gorm_mysql

import (
	"log/slog"

	"github.com/kordar/godb"
	"gorm.io/gorm"
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
	mysqlConfig.Logger = newSlogGormLogger(dbLogLevel)
	return &mysqlConfig
}

// AddMysqlInstances 批量添加mysql句柄
func AddMysqlInstances(dbs map[string]map[string]string) {
	for db, cfg := range dbs {
		err := AddMysqlInstance(db, cfg)
		if err != nil {
			slog.Warn("add mysql instance failed", "err", err)
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
