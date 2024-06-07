package mysql

import (
	"fmt"
	"trainee3/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IMysqlDB interface {
	GetDB(dbName DBName) *gorm.DB
}

type mysqlDB struct {
	db map[DBName]*gorm.DB
}

type DBName string

const (
	Trainee3   DBName = "trainee3"
	Trainee3_1 DBName = "trainee3_1"
)

// Mysql版
func New(cfg *config.Config) (IMysqlDB, error) {
	var logMode logger.LogLevel
	if cfg.Other.Mode == "PROUDCTION" {
		logMode = logger.Silent
	} else {
		logMode = logger.Info
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		cfg.MysqlConfig.UserName, cfg.MysqlConfig.Password, cfg.MysqlConfig.Addr, cfg.MysqlConfig.Port, cfg.MysqlConfig.Database)), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	return mysqlDB{
		map[DBName]*gorm.DB{
			Trainee3:   db,
			Trainee3_1: db,
		},
	}, nil
}

func (mdb mysqlDB) GetDB(dbName DBName) *gorm.DB {
	return mdb.db[dbName]
}
