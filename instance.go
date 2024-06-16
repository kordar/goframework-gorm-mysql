package goframework_gorm_mysql

import (
	"fmt"
	"github.com/kordar/gocfg"
	log "github.com/kordar/gologger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormConnIns struct {
	name string
	ins  *gorm.DB
}

func NewGormConnIns(name string, config *gorm.Config) *GormConnIns {
	cfg := gocfg.GetSection(name)
	return NewGormConnInsWithConfig(name, cfg, config)
}

func NewGormConnInsWithConfig(name string, cfg map[string]string, config *gorm.Config) *GormConnIns {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg["user"], cfg["password"], cfg["host"], cfg["port"], cfg["db"], "charset="+cfg["charset"]+"&parseTime=true")
	return NewGormConnInsWithDsn(name, dsn, config)
}

func NewGormConnInsWithDsn(name string, dsn string, config *gorm.Config) *GormConnIns {
	ins, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Errorf("初始化mysql失败,dsn=%s,err=%v", dsn, err)
		return nil
	}
	return &GormConnIns{name: name, ins: ins}
}

func (c GormConnIns) GetName() string {
	return c.name
}

func (c GormConnIns) GetInstance() interface{} {
	return c.ins
}

func (c GormConnIns) Close() error {
	if s, err := c.ins.DB(); err == nil {
		return s.Close()
	} else {
		return err
	}
}
