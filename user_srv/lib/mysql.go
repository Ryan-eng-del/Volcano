package lib

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"volcano.user_srv/config"
	"volcano.user_srv/internal"
)

type Mysql struct {
	conf config.MysqlConfig

}

var DB *gorm.DB

func NewMysql(conf config.MysqlConfig) *Mysql {
	return &Mysql{conf}
}

func (m *Mysql) Init() error {
	c := m.conf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.TimeLocation)
	log.Println(dsn)
	zap.S().Infof("connect dsn: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: &internal.DefaultGormLogger,
	}); 
	
	if  err != nil {
		zap.S().Errorf("lib.mysql.Init.Open: %s", err.Error())
		return err
	}

	DB = db
	dbpool, err := db.DB()

	if err != nil {
		zap.S().Errorf("lib.mysql.Init.DB: %s", err.Error())
		return err
	}

	dbpool.SetMaxOpenConns(c.MaxOpenConn)
	dbpool.SetMaxIdleConns(c.MaxIdleConn)
	dbpool.SetConnMaxLifetime(time.Duration(c.MaxCoonLifeTime) * time.Second)

	err = dbpool.Ping()
	if err != nil {
		zap.S().Errorf("lib.mysql.Init.DB: %s", err.Error())
		return err
	}

	return nil
}