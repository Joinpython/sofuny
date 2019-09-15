package models

type Comment struct {
	BaseModel
	ArticleID int    `json:"article_id" gorm:"type:Int;not null"`
	Name      string `json:"name" gorm:"type:varchar(10);unique;not null"`
	Email     string `json:"email" gorm:"type:varchar(40);null"`
	Content   string `json:"content" gorm:"type:varchar(100);not null"`
	Like      int    `json:"like" gorm:"type:int;default(0)"`
	DisLike   int    `json:"dislike" gorm:"type:int;default(0)"`
}
