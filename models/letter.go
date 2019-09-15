package models

import "time"

// 信件
type Letter struct {
	BaseModel
	Name       string    `json:"name" gorm:"varchar(20);not null"`
	Email      string    `json:"email" gorm:"varchar(40);not null"`
	SendAt     time.Time `json:"send_at"`
	Content    string    `json:"content" gorm:"varchar(200);not null"`
	SendStatus bool      `json:"send_status"`
	Browser    string    `json:"browser" gorm:"varchar(15);not null"`
	Device     string    `json:"device" gorm:"varchar(15);not null"`
	LikeCounts int       `json:"like_counts" gorm:"type:Int;default 0"`
}
