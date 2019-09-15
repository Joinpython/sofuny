package models

type Article struct {
	BaseModel
	UserID        int    `json:"userID" gorm:"type:Int;not null"`
	Title         string `json:"title" gorm:"type:varchar(50);unique;not null"`
	Category      int    `json:"category" gorm:"not null"`
	Content       string `json:"content" gorm:"type:longtext;not null"`
	Images        string `json:"images" gorm:"text;not null"`
	MediaUrl      string `json:"mediaUrl" gorm:"varchar(200);not null"`
	ViewCounts    int    `json:"viewCounts" gorm:"type:Int;default(0)"`
	CommentCounts int    `json:"commentCounts" gorm:"type:Int;default(0)"`
	LikeCounts    int    `json:"likeCounts" gorm:"type:Int;default(0)"`
}
