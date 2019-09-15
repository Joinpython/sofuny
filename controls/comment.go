package controls

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"sofuny/models"
	"sofuny/utils"
	"time"
)

// 创建 评论
func CreateComment(ctx echo.Context) error {
	var comment models.Comment
	if err := ctx.Bind(&comment); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	var lastComment models.Comment
	comment.Uuid = xid.New().String()
	db.Find(&lastComment)
	comment.ID = lastComment.ID + 1
	if err := db.Create(&comment).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	} else {
		var article models.Article
		db.Where("id=?", comment.ArticleID).Find(&article)
		article.CommentCounts += 1
		db.Save(&article)
		return ctx.JSON(200, utils.Response{
			StatusCode: 200,
			Msg:        "success",
			Data:       comment,
			Time:       time.Now().Local(),
		})
	}
}

// 查找评论
func FindComment(ctx echo.Context) error {
	article_id := ctx.QueryParam("article_id")
	var commentList []models.Comment
	if err := db.Where("article_id=?", article_id).Find(&commentList).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       commentList,
		Time:       time.Now().Local(),
	})
}

// 评论喜欢
type CommentLikes struct {
	ID int `json:"id"`
}

func LikeComment(ctx echo.Context) error {
	var commentLike CommentLikes
	var comment models.Comment
	if err := ctx.Bind(&commentLike); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Limit(1).Where("id=?", commentLike.ID).First(&comment).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	comment.Like += 1
	db.Save(&comment)
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       map[string]int{"newCommentLike": comment.Like},
		Time:       time.Now().Local(),
	})
}

// 评论不喜欢
func DislikeComment(ctx echo.Context) error {
	var commentDisLike CommentLikes
	var comment models.Comment
	if err := ctx.Bind(&commentDisLike); err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	if err := db.Limit(1).Where("id=?", commentDisLike.ID).First(&comment).Error; err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg:        err,
			Time:       time.Now().Local(),
		})
	}
	comment.DisLike += 1
	db.Save(&comment)
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg:        "success",
		Data:       map[string]int{"newCommentDislike": comment.DisLike},
		Time:       time.Now().Local(),
	})
}
