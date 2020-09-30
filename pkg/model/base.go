package model

import (
	"fmt"
	"time"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gin_gateway/pkg/conf"
)

var db *gorm.DB

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}

func SetupDB() {
	var err error
	fmt.Println(conf.GlobalConfig.Db)
	db, err = gorm.Open(conf.GlobalConfig.Db.DbType,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.GlobalConfig.Db.DbUser,
			conf.GlobalConfig.Db.DbPassword,
			conf.GlobalConfig.Db.DbHost,
			conf.GlobalConfig.Db.DbPort,
			conf.GlobalConfig.Db.DbName))
	if err != nil {
		log.Fatalf("models.SetupAdmin err: %s", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}