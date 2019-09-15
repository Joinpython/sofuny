package controls

import (
	"github.com/labstack/echo/v4"
	"sofuny/models"
	"sofuny/utils"
	"time"
)

// 分类 文章
type Cate struct { Cate int `json:"cate"` }
func FindCategory(ctx echo.Context) error {
	var results []models.Article
	var cate Cate
	if err := ctx.Bind(&cate);err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg: err,
			Time: time.Now().Local(),
		})
	}
	if err := db.Where("category=?", cate.Cate).Find(&results).Error;err != nil {
		return ctx.JSON(200, utils.Response{
			StatusCode: 201,
			Msg: err,
			Time: time.Now().Local(),
		})
	}
	return ctx.JSON(200, utils.Response{
		StatusCode: 200,
		Msg: "success",
		Data: results,
		Time: time.Now().Local(),
	})
}
