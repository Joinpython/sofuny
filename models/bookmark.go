package models

// 大分类
type BigCategory struct {
	BaseModel
	Name              string             `json:"name" gorm:"type:varchar(10);unique;not null"`
	BookmarkCategorys []BookmarkCategory `json:"bookmark_categorys" gorm:"foreignkey:BigID;null"`
}

// 小分类
type BookmarkCategory struct {
	BaseModel
	BigID     int        `json:"big_id" gorm:"type:Int;not null"`
	Name      string     `json:"name" gorm:"type:varchar(10);unique;not null"`
	Bookmarks []Bookmark `json:"bookmarks" gorm:"foreignkey:CateID;null"`
}

// 书签
type Bookmark struct {
	BaseModel
	UserID   int    `json:"user_id" gorm:"type:Int;not null"`
	CateID   int    `json:"cate_id" gorm:"type:Int;not null"`
	Name     string `json:"name" gorm:"varchar(50);not null"`
	Url      string `json:"url" gorm:"type:varchar(200);not null"`
	Category string `json:"category" gorm:"varchar(20);not null"`
	Abstract string `json:"abstract" gorm:"varchar(200); not null"`
}
