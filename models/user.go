package models

// 用户
type User struct {
	BaseModel
	Name      string     `json:"name" gorm:"type:varchar(10);unique;not null"`
	Email     string     `json:"email" gorm:"type:varchar(30);unique;not null"`
	Password  string     `json:"password" gorm:"type:varchar(200);not null"`
	IsSuper   bool       `json:"is_super" gorm:"default false"`
	Bookmarks []Bookmark `json:"bookmarks" gorm:"foreignkey:UserID;null"`
	Articles  []Article  `json:"bookmarks" gorm:"foreignkey:UserID;null"`
}
