package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sofuny/config"
	"time"
)

// base models
type BaseModel struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status"`
	Uuid      string    `json:"uuid" gorm:"Type:varchar(40);unique;not null"`
}

var database = config.Config().Database

// 连接数据库
func Connection() *gorm.DB {
	str := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", database.User, database.Password, database.Name)
	db, err := gorm.Open(database.Driver, str)
	if err != nil {
		fmt.Println("连接--failed!")
	}
	db.DB().SetMaxIdleConns(1000)
	db.DB().SetMaxOpenConns(1000)
	// 启用Logger，显示详细日志
	db.LogMode(false)
	// SetConnMaxLifetime设置连接可以重用的最长时间。
	//db.DB().SetConnMaxLifetime(time.Hour)
	fmt.Println("连接--success!")
	return db
}
